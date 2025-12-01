<!-- src/components/common/RecordSidebar.vue -->
<template>
	<!-- 记录侧边栏 -->
	<div class="record-sidebar" :class="{ active: modelValue }">
		<div class="sidebar-header">
			<h3>{{ title }}</h3>
			<button class="close-btn" @click="closeSidebar">×</button>
		</div>
		<div class="record-list">
			<div v-if="records.length === 0" class="empty-records">
				{{ emptyText }}
			</div>
			<div v-else class="records-table">
				<div v-for="(record, index) in records" :key="index" class="record-item">
					<div class="file-name">{{ record.fileName }}</div>
					<div class="record-time">{{ record.time }}</div>
				</div>
			</div>
		</div>
	</div>

	<!-- 侧边栏遮罩层 -->
	<div class="sidebar-overlay" v-if="modelValue" @click="closeSidebar"></div>
</template>

<script setup>
	const props = defineProps({
		modelValue: {
			type: Boolean,
			default: false
		},
		title: {
			type: String,
			default: "记录"
		},
		emptyText: {
			type: String,
			default: "暂无记录"
		},
		records: {
			type: Array,
			default: () => []
		}
	})

	const emit = defineEmits(['update:modelValue'])

	const closeSidebar = () => {
		emit('update:modelValue', false)
	}
</script>

<style scoped>
	/* 记录侧边栏样式 */
	.record-sidebar {
		position: fixed;
		top: 0;
		right: -400px;
		width: 350px;
		height: 100vh;
		background: white;
		box-shadow: -2px 0 12px rgba(0, 0, 0, 0.15);
		transition: right 0.3s ease;
		z-index: 1000;
		display: flex;
		flex-direction: column;
	}

	.record-sidebar.active {
		right: 0;
	}

	.sidebar-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 20px;
		border-bottom: 1px solid #f0f0f0;
		background: #fafafa;
	}

	.sidebar-header h3 {
		margin: 0;
		font-size: 18px;
		font-weight: 600;
		color: #333;
	}

	.close-btn {
		background: none;
		border: none;
		font-size: 24px;
		cursor: pointer;
		color: #666;
		width: 30px;
		height: 30px;
		display: flex;
		align-items: center;
		justify-content: center;
		border-radius: 50%;
	}

	.close-btn:hover {
		background: #f0f0f0;
		color: #333;
	}

	.record-list {
		flex: 1;
		overflow: hidden;
		display: flex;
		flex-direction: column;
	}

	.empty-records {
		text-align: center;
		color: #999;
		padding: 40px 20px;
		font-size: 14px;
	}

	.records-table {
		flex: 1;
		overflow-y: auto;
		padding: 0;
	}

	.record-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 16px 20px;
		border-bottom: 1px solid #f8f8f8;
		transition: background-color 0.2s;
	}

	.record-item:hover {
		background-color: #f9f9f9;
	}

	.record-item:last-child {
		border-bottom: none;
	}

	.file-name {
		flex: 1;
		font-size: 14px;
		color: #333;
		word-break: break-all;
		padding-right: 12px;
	}

	.record-time {
		font-size: 12px;
		color: #666;
		white-space: nowrap;
	}

	/* 侧边栏遮罩层 */
	.sidebar-overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background: rgba(0, 0, 0, 0.5);
		z-index: 999;
	}

	/* 移动端适配 */
	@media (max-width: 768px) {
		.record-sidebar {
			width: 100%;
			right: -100%;
		}
	}
</style>