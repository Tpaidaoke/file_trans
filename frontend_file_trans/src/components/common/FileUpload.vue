<!-- src/components/common/FileUpload.vue -->
<template>
	<div class="upload-area" @click="triggerFileInput"
		:class="{ uploading: isUploading, 'has-files': selectedFiles.length > 0 }">

		<!-- 四条边框进度条 -->
		<div class="progress-border-container">
			<div class="border-progress top" :style="{ transform: `scaleX(${getBorderProgress(0)})` }"></div>
			<div class="border-progress right" :style="{ transform: `scaleY(${getBorderProgress(1)})` }"></div>
			<div class="border-progress bottom" :style="{ transform: `scaleX(${getBorderProgress(2)})` }"></div>
			<div class="border-progress left" :style="{ transform: `scaleY(${getBorderProgress(3)})` }"></div>
		</div>

		<!-- 上传提示（有文件时隐藏） -->
		<div class="upload-placeholder" v-if="selectedFiles.length === 0 && !isUploading">
			点击或拖拽文件到此处上传
		</div>
		<div class="upload-hint" v-if="selectedFiles.length === 0 && !isUploading">
			支持待编辑页面，最大 200GB
		</div>

		<input type="file" ref="fileInput" @change="handleFileUpload" style="display: none" multiple>

		<!-- 显示已选择的文件 -->
		<div v-if="selectedFiles.length > 0" class="selected-files">
			<!-- 当前上传文件状态 -->
			<div v-if="isUploading" class="uploading-status">
				正在上传: {{ uploadingFileName }} ({{ uploadProgress }}%)
			</div>

			<!-- 文件列表 -->
			<div v-for="(file, index) in selectedFiles" :key="index" class="file-item"
				:class="{ 'uploading': file.uploading, 'completed': file.completed }">
				<div class="file-info">
					<span class="file-name">{{ file.name }}</span>
					<span class="file-size">({{ formatFileSize(file.size) }})</span>
				</div>
				<div class="file-status">
					<!-- 上传进度 -->
					<div v-if="file.uploading" class="file-progress">
						<div class="progress-bar">
							<div class="progress-fill" :style="{ width: `${file.progress}%` }"></div>
						</div>
						<span class="progress-text">{{ file.progress }}%</span>
					</div>
					<!-- 上传完成 -->
					<span v-else-if="file.completed" class="status-completed">✓ 完成</span>
					<!-- 等待上传 -->
					<span v-else class="status-waiting">等待上传</span>
					<button @click.stop="removeFile(index)" class="remove-btn">×</button>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup>
	import {
		ref,
		computed
	} from 'vue'
	import {
		useFileUpload
	} from '@/composables/useFileUpload'

	const emit = defineEmits(['file-uploaded'])

	const {
		fileInput,
		selectedFiles,
		isUploading,
		uploadProgress,
		uploadingFileName,
		triggerFileInput,
		handleFileUpload,
		removeFile,
		simulateUploadProgress
	} = useFileUpload(emit)

	// 计算四条边框的进度
	const getBorderProgress = (borderIndex) => {
		if (!isUploading.value) return 0

		const progress = uploadProgress.value
		const segment = 25 // 每条边占25%
		const startProgress = borderIndex * segment
		const endProgress = (borderIndex + 1) * segment

		if (progress <= startProgress) return 0
		if (progress >= endProgress) return 1

		// 当前边内的进度 (0-1)
		return (progress - startProgress) / segment
	}

	// 格式化文件大小
	const formatFileSize = (bytes) => {
		if (bytes === 0) return '0 Bytes'
		const k = 1024
		const sizes = ['Bytes', 'KB', 'MB', 'GB']
		const i = Math.floor(Math.log(bytes) / Math.log(k))
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
	}
</script>

<style scoped>
	.upload-area {
		border: 2px dashed #d9d9d9;
		padding: 40px 20px;
		border-radius: 8px;
		text-align: center;
		transition: all 0.3s;
		width: 100%;
		box-sizing: border-box;
		position: relative;
		overflow: hidden;
	}

	.upload-area:hover {
		border-color: #1890ff;
		background: #f0f8ff;
	}

	.upload-area.uploading {
		border-color: transparent;
	}

	.upload-area.has-files {
		padding: 20px;
	}

	/* 四条边框进度条容器 */
	.progress-border-container {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		pointer-events: none;
		border-radius: 8px;
	}

	/* 通用边框进度条样式 */
	.border-progress {
		position: absolute;
		background: linear-gradient(90deg, #667eea, #764ba2);
		transition: transform 0.3s ease;
		transform-origin: left center;
		box-shadow: 0 0 8px rgba(102, 126, 234, 0.6);
		animation: pulse-glow 2s ease-in-out infinite;
	}

	/* 上边框 */
	.border-progress.top {
		top: 0;
		left: 0;
		width: 100%;
		height: 3px;
		transform-origin: left center;
	}

	/* 右边框 */
	.border-progress.right {
		top: 0;
		right: 0;
		width: 3px;
		height: 100%;
		transform-origin: center top;
	}

	/* 下边框 */
	.border-progress.bottom {
		bottom: 0;
		right: 0;
		width: 100%;
		height: 3px;
		transform-origin: right center;
	}

	/* 左边框 */
	.border-progress.left {
		bottom: 0;
		left: 0;
		width: 3px;
		height: 100%;
		transform-origin: center bottom;
	}

	.upload-placeholder {
		font-size: 16px;
		color: #666;
		margin-bottom: 8px;
	}

	.upload-hint {
		font-size: 14px;
		color: #999;
	}

	/* 文件列表样式 */
	.selected-files {
		width: 100%;
		text-align: left;
	}

	.uploading-status {
		background: #e6f7ff;
		padding: 8px 12px;
		border-radius: 4px;
		margin-bottom: 12px;
		color: #1890ff;
		font-size: 14px;
	}

	.file-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 12px;
		border: 1px solid #f0f0f0;
		border-radius: 6px;
		margin-bottom: 8px;
		background: white;
		transition: all 0.3s;
	}

	.file-item:hover {
		box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	}

	.file-item.uploading {
		border-left: 3px solid #1890ff;
	}

	.file-item.completed {
		border-left: 3px solid #52c41a;
	}

	.file-info {
		flex: 1;
	}

	.file-name {
		font-weight: 500;
		color: #333;
	}

	.file-size {
		color: #666;
		font-size: 12px;
		margin-left: 8px;
	}

	.file-status {
		display: flex;
		align-items: center;
		gap: 12px;
	}

	.file-progress {
		display: flex;
		align-items: center;
		gap: 8px;
	}

	.progress-bar {
		width: 60px;
		height: 6px;
		background: #f0f0f0;
		border-radius: 3px;
		overflow: hidden;
	}

	.progress-fill {
		height: 100%;
		background: #1890ff;
		transition: width 0.3s ease;
	}

	.progress-text {
		font-size: 12px;
		color: #1890ff;
		min-width: 30px;
	}

	.status-completed {
		color: #52c41a;
		font-size: 12px;
	}

	.status-waiting {
		color: #666;
		font-size: 12px;
	}

	.remove-btn {
		background: #ff4d4f;
		color: white;
		border: none;
		border-radius: 50%;
		width: 24px;
		height: 24px;
		display: flex;
		align-items: center;
		justify-content: center;
		cursor: pointer;
		font-size: 16px;
		transition: all 0.3s;
	}

	.remove-btn:hover {
		background: #ff7875;
	}

	@keyframes pulse-glow {

		0%,
		100% {
			box-shadow: 0 0 8px rgba(102, 126, 234, 0.6);
		}

		50% {
			box-shadow: 0 0 12px rgba(102, 126, 234, 0.9), 0 0 16px rgba(102, 126, 234, 0.4);
		}
	}
</style>