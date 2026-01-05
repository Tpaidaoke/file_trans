<template>
	<!-- 记录侧边栏 -->
	<div class="record-sidebar" :class="{ active: modelValue }">
		<div class="sidebar-header">
			<h3>{{ title }}</h3>
			<button class="close-btn" @click="closeSidebar">×</button>
		</div>
		<div class="record-list">
			<div v-if="records.length === 0" class="empty-records">
				<!-- 增加空状态图标（文字图标，无需引入资源） -->
				<div class="empty-icon">📄</div>
				<div class="empty-text">{{ emptyText }}</div>
			</div>
			<div v-else class="records-table">
				<div v-for="(record, index) in records" :key="index" class="record-item">
					<!-- 序号 -->
					<div class="record-index">{{ index+1 }}</div>
					<!-- 文件名 -->
					<div class="file-name">{{ record.FileName }}</div>
					<!-- 时间 -->
					<div class="record-time">{{ dayjs(record.ReceiveAt).format('YYYY-MM-DD HH:mm:ss') }}</div>
				</div>
			</div>
		</div>
	</div>

	<!-- 侧边栏遮罩层 -->
	<div class="sidebar-overlay" v-if="modelValue" @click="closeSidebar"></div>
</template>

<script setup>	
	import dayjs from 'dayjs';

	const props = defineProps({
		modelValue: { type: Boolean, default: false },
		title: { type: String, default: "记录" },
		emptyText: { type: String, default: "暂无记录" },
		records: { type: Array, default: () => [] }
	})

	const emit = defineEmits(['update:modelValue'])

	const closeSidebar = () => {
		emit('update:modelValue', false)
	}
</script>

<style scoped>
	/* 核心优化：整体风格更柔和、层次更清晰 */
	:root {
		--color-primary: #409EFF; /* 主色调（柔和蓝色） */
		--color-text-main: #303133; /* 主文字色 */
		--color-text-secondary: #909399; /* 次要文字色 */
		--color-bg-light: #F5F7FA; /* 浅背景色 */
		--color-border: #E4E7ED; /* 边框色 */
		--shadow: 0 2px 12px rgba(0, 0, 0, 0.08); /* 柔和阴影 */
	}

	/* 记录侧边栏容器 */
	.record-sidebar {
		position: fixed;
		top: 0;
		right: -400px;
		width: 380px; /* 适度加宽，提升排版空间 */
		height: 100vh;
		background: #FFFFFF;
		box-shadow: var(--shadow);
		transition: right 0.3s cubic-bezier(0.25, 0.8, 0.25, 1); /* 更丝滑的动画 */
		z-index: 1000;
		display: flex;
		flex-direction: column;
		border-left: 1px solid var(--color-border);
	}

	.record-sidebar.active {
		right: 0;
	}

	/* 侧边栏头部 */
	.sidebar-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 16px 24px;
		background: var(--color-bg-light);
		border-bottom: 1px solid var(--color-border);
	}

	.sidebar-header h3 {
		margin: 0;
		font-size: 16px;
		font-weight: 600;
		color: var(--color-text-main);
	}

	/* 关闭按钮优化 */
	.close-btn {
		background: transparent;
		border: none;
		font-size: 20px;
		cursor: pointer;
		color: var(--color-text-secondary);
		width: 32px;
		height: 32px;
		display: flex;
		align-items: center;
		justify-content: center;
		border-radius: 50%;
		transition: all 0.2s;
	}

	.close-btn:hover {
		background: #E8E8E8;
		color: var(--color-text-main);
	}

	/* 记录列表容器 */
	.record-list {
		flex: 1;
		overflow: hidden;
		display: flex;
		flex-direction: column;
	}

	/* 空状态优化（增加图标、调整间距） */
	.empty-records {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 60px 20px;
		color: var(--color-text-secondary);
	}

	.empty-icon {
		font-size: 40px;
		margin-bottom: 16px;
		opacity: 0.6;
	}

	.empty-text {
		font-size: 14px;
	}

	/* 记录表格容器（美化滚动条） */
	.records-table {
		flex: 1;
		overflow-y: auto;
		padding: 8px 0;
		/* 自定义滚动条 */
		&::-webkit-scrollbar {
			width: 6px;
			height: 6px;
		}
		&::-webkit-scrollbar-thumb {
			background-color: rgba(144, 147, 153, 0.3);
			border-radius: 3px;
		}
		&::-webkit-scrollbar-track {
			background-color: transparent;
		}
	}

	/* 记录项（核心布局优化） */
	.record-item {
		display: flex;
		align-items: center;
		padding: 14px 24px;
		border-bottom: 1px solid #F2F2F2;
		transition: all 0.2s;
		border-radius: 4px;
		margin: 0 8px;
	}

	/* hover 效果优化（增加轻微悬浮感） */
	.record-item:hover {
		background-color: var(--color-bg-light);
		transform: translateX(2px);
	}

	/* 序号（固定宽度，淡色区分） */
	.record-index {
		width: 24px;
		height: 24px;
		line-height: 24px;
		text-align: center;
		font-size: 12px;
		color: var(--color-primary);
		background-color: rgba(64, 158, 255, 0.1);
		border-radius: 4px;
		margin-right: 12px;
	}

	/* 文件名（占主要空间，优化换行） */
	.file-name {
		flex: 1;
		font-size: 14px;
		color: var(--color-text-main);
		word-break: break-all;
		line-height: 24px;
	}

	/* 时间（右对齐，次要文字色） */
	.record-time {
		font-size: 12px;
		color: var(--color-text-secondary);
		white-space: nowrap;
		margin-left: 16px;
		line-height: 24px;
	}

	/* 侧边栏遮罩层（更柔和的透明度） */
	.sidebar-overlay {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background: rgba(0, 0, 0, 0.3);
		z-index: 999;
		backdrop-filter: blur(2px); /* 增加模糊效果，提升质感 */
	}

	/* 移动端适配（保持全屏宽度） */
	@media (max-width: 768px) {
		.record-sidebar {
			width: 100%;
			right: -100%;
		}
	}
</style>