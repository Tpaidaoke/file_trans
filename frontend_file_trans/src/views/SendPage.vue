<!-- src/views/SendPage.vue -->
<template>
	<div class="send-page">
		<!-- 页面头部 -->
		<CommonHeader title="发文件" />

		<!-- 主要内容区域 -->
		<div class="content">
			<SendForm @show-records="showRecordSidebar = true" @submit="handleFormSubmit" />
		</div>

		<!-- 发件记录侧边栏 -->
		<RecordSidebar title="发件记录" empty-text="暂无发件记录" v-model="showRecordSidebar" :records="todaySendRecords" />
	</div>
</template>

<script setup>
	import {
		ref
	} from 'vue'
	import CommonHeader from '@/components/common/Header.vue'
	import SendForm from '@/components/send/SendForm.vue'
	import RecordSidebar from '@/components/common/RecordSidebar.vue'
	import {useSendRecords} from '@/composables/useSendRecords'

	const {todaySendRecords} = useSendRecords()
	// 发件记录侧边栏显示状态
	const showRecordSidebar = ref(false)

	// 处理表单提交
	const handleFormSubmit = (submitData) => {
		console.log('表单提交数据:', submitData)

		// 这里可以添加实际的提交逻辑
		// 比如调用API发送数据

		// 模拟提交成功
		alert('发送成功！')
	}
</script>

<style scoped>
	.send-page {
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
	}

	/* 响应式设计 */
	@media (max-width: 768px) {
		.send-page {
			padding: 16px;
		}

		.content {
			padding: 16px;
		}
	}
</style>