<template>
	<div class="receive-page">
		<!-- 页面头部 -->
		<CommonHeader title="取文件" />

		<!-- 主要内容区域 -->
		<div class="content">
			<RecForm @show-records="showRecordSidebar = true" @submit="handleFormSubmit" />
		</div>

		<!-- 发件记录侧边栏 -->
		<RecordSidebar title="取件记录" empty-text="暂无取件记录" v-model="showRecordSidebar" :records="todayPickupRecords"/>
	</div>
</template>

<script setup>
	import {
		ref
	} from 'vue';
	import CommonHeader from '@/components/common/Header.vue'
	import RecForm from '@/components/receive/ReceiveForm.vue'
	import RecordSidebar from '@/components/common/RecordSidebar.vue'
	import { usePickupRecords } from '@/composables/usePickupRecords'

	const { pickupRecords } = usePickupRecords()
	// 发件记录侧边栏显示状态
	const showRecordSidebar = ref(false)
	
	// 计算今日取件记录
	const todayPickupRecords = ref([])

	// 处理表单提交
	const handlePickupSuccess = (fileName) => {
	  console.log('取件成功:', fileName)
	}
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
	}
</style>