<!-- src/components/common/FileUpload.vue -->
<template>
	<div class="upload-area" @click="triggerFileInput" ref="dropAreaRef"
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
			支持待编辑页面，最大 1GB，最多 {{ MAX_FILES }} 个文件
		</div>

		<input type="file" ref="fileInput" @change="handleFileUpload" style="display: none" multiple>

		<!-- 显示已选择的文件 -->
		<div v-if="selectedFiles.length > 0" class="selected-files">
			<!-- 文件列表头部：文件数量统计 -->
			<div class="files-header" v-if="selectedFiles.length > 0">
				<div class="files-count">
					已选择 {{ selectedFiles.length }} 个文件
					<span v-if="isUploading">(上传中: {{ getUploadingCount() }}/{{ selectedFiles.length }})</span>
				</div>
				<div class="files-total-size">
					总大小: {{ formatFileSize(getTotalSize()) }}
				</div>
			</div>

			<!-- 文件列表 -->
			<div v-for="(file, index) in selectedFiles" :key="index" class="file-item"
				:class="{ 'uploading': file.uploading, 'completed': file.completed }">
				<div class="file-info">
					<span class="file-name">{{ file.name }}</span>
					<!-- 确认文件大小显示正确 -->
					<span class="file-size">({{ formatFileSize(file.size) }})</span>
					<!-- 显示单个文件上传进度 -->
					<span class="file-progress" v-if="isUploading">
						{{ file.progress }}%
					</span>
				</div>
				<button @click.stop="removeFile(index)" class="remove-btn" :disabled="isUploading">×</button>
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
	import {
		formatFileSize
	} from '@/utils/fileUtils'

	const emit = defineEmits(['file-uploaded'])

	// 拖拽区域引用
	const dropAreaRef = ref(null)

	const {
		fileInput,
		selectedFiles,
		isUploading,
		uploadProgress,
		uploadingFileName,
		triggerFileInput,
		handleFileUpload,
		removeFile,
		simulateFileUpload,
		clearFileList,
		MAX_FILES
	} = useFileUpload(emit, dropAreaRef)

	// 计算上传中的文件数量
	const getUploadingCount = () => {
		return selectedFiles.value.filter(file => file.uploading || file.completed).length
	}

	// 计算总文件大小
	const getTotalSize = () => {
		return selectedFiles.value.reduce((total, file) => total + file.size, 0)
	}

	// 适配进度条的transform-origin，重新计算四条边的进度
	const getBorderProgress = (borderIndex) => {
		if (!isUploading.value) return 0

		const progress = uploadProgress.value / 100 // 转为0-1的比例
		const segment = 0.25 // 每条边占25%（对应0-0.25）

		// 四个边的进度计算逻辑（适配各自的transform-origin）
		switch (borderIndex) {
			case 0: // 上边框（transform-origin: left center）：0 → 1
				return progress >= segment ? 1 : progress / segment
			case 1: // 右边框（transform-origin: center top）：0 → 1
				return progress <= segment ? 0 : (progress - segment) / segment
			case 2: // 下边框（transform-origin: right center）：0 → 1
				return progress <= 2 * segment ? 0 : (progress - 2 * segment) / segment
			case 3: // 左边框（transform-origin: center bottom）：0 → 1
				return progress <= 3 * segment ? 0 : (progress - 3 * segment) / segment
			default:
				return 0
		}
	}

	defineExpose({
		clearFileList // 暴露清空方法
	})
</script>

<style scoped>
	/* 原有样式保持不变，添加新样式 */
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
		box-shadow: 0 0 8px rgba(102, 126, 234, 0.6);
		animation: pulse-glow 2s ease-in-out infinite;
	}

	/* 上边框：transform-origin: left center */
	.border-progress.top {
		top: 0;
		left: 0;
		width: 100%;
		height: 3px;
		transform-origin: left center;
	}

	/* 右边框：transform-origin: center top */
	.border-progress.right {
		top: 0;
		right: 0;
		width: 3px;
		height: 100%;
		transform-origin: center top;
	}

	/* 下边框：transform-origin: right center */
	.border-progress.bottom {
		bottom: 0;
		right: 0;
		width: 100%;
		height: 3px;
		transform-origin: right center;
	}

	/* 左边框：transform-origin: center bottom */
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

	/* 文件列表头部样式 */
	.files-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 12px;
		padding: 8px 12px;
		background-color: #f8f9fa;
		border-radius: 6px;
		font-size: 14px;
		color: #495057;
	}

	.files-count {
		font-weight: 500;
	}

	.files-total-size {
		color: #6c757d;
	}

	/* 文件项中的进度显示 */
	.file-progress {
		margin-left: 8px;
		font-size: 12px;
		color: #1890ff;
		font-weight: 500;
		background: rgba(24, 144, 255, 0.1);
		padding: 2px 6px;
		border-radius: 10px;
	}

	/* 禁用状态的删除按钮 */
	.remove-btn:disabled {
		background: #d9d9d9;
		cursor: not-allowed;
		opacity: 0.6;
	}

	.remove-btn:disabled:hover {
		background: #d9d9d9;
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