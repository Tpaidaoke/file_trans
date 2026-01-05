<template>
	<div class="receive-page">
		<!-- 页面头部 -->
		<CommonHeader title="取文件" />

		<!-- 主要内容区域 -->
		<div class="content">
			<!-- 接收表单组件：绑定提交事件和加载状态 -->
			<RecForm @show-records="handleShowRecords" @pickup-success="handleFormSubmit" :loading="isLoading" />
		</div>

		<!-- 发件记录侧边栏 -->
		<RecordSidebar title="取件记录" empty-text="暂无取件记录" v-model="showRecordSidebar" :records="todayPickupRecords" />
	</div>
</template>

<script setup>
	import {
		ref,
		onMounted
	} from 'vue';
	import axios from 'axios';
	import CommonHeader from '@/components/common/Header.vue';
	import RecForm from '@/components/receive/ReceiveForm.vue';
	import RecordSidebar from '@/components/common/RecordSidebar.vue';

	// 状态管理
	const showRecordSidebar = ref(false); // 侧边栏显示状态
	const isLoading = ref(false); // 表单提交加载状态
	const todayPickupRecords = ref([]); // 今日取件记录

	
	// 从后端获取取件记录
	const fetchPickupRecords = async () => {
		try {
			const response = await axios.get('/api/v1/receiveRecords')
			todayPickupRecords.value = response.data.records || []
		} catch (error) {
			console.error('获取取件记录失败', error)
		}
	}
	
	// 处理点击取件记录的方法
	const handleShowRecords = async () => {
		// 等待获取最新的取件记录返回后显示
		await fetchPickupRecords()
		console.log('取件记录数据:', todayPickupRecords)
		showRecordSidebar.value = true
	}

	// 定义handlePickupSuccess方法
	const handlePickupSuccess = async (fileInfo) => {
		// 调用后端接口更新文件状态（下载完成）
		try {
			console.log("fileInfo", fileInfo, fileInfo.fileUuid)
			await axios.post('/api/v1/updateFileStatus', {
				fileUuid: fileInfo.fileUuid
			});
			console.log('文件状态已更新为已下载');
		} catch (error) {
			console.error('更新文件状态失败:', error);
		}
	};

	// 处理表单提交（取件操作）
	const handleFormSubmit = (fileData) => {
		const {
			fileName,
			fileUuid,
			fileSize,
			expired,
			fileDownloadUrl
		} = fileData;
		console.log("返回的数据:", fileData)

		if (expired) {
			console.log('文件已过期', 'error');
			return;
		}

		// 调用handlePickupSuccess并传递完整信息
		handlePickupSuccess({
			fileName,
			fileSize,
			fileDownloadUrl,
			fileUuid, // 传递fileUuid用于状态更新
			pickupTime: new Date().toLocaleString()
		});

		console.log('取件成功，正在下载文件', 'success');
		// 触发文件下载
		window.open(fileDownloadUrl, '_blank');
	};
</script>

<style>
	.receive-page {
		max-width: 600px;
		margin: 0 auto;
		padding: 20px;
		font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
		color: #333;
		position: relative;
	}

	.content {
		background: #fff;
		border-radius: 12px;
		box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
		padding: 24px;
		margin-top: 20px;
	}
</style>