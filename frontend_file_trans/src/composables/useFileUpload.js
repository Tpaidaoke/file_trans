// src/composables/useFileUpload.js
import {
	ref,
	onMounted,
	onUnmounted
} from 'vue'

export function useFileUpload(emit, dropAreaRef = null) {
	const fileInput = ref(null)
	const selectedFiles = ref([])
	const isUploading = ref(false) // 全局上传状态（控制进度条显示）
	const uploadProgress = ref(0) // 全局上传进度（控制四条边框）
	const uploadingFileName = ref('')
	let progressTimer = null;

	const MAX_FILES = 10

	const triggerFileInput = () => {
		fileInput.value.click()
	}

	const handleFileUpload = async (event) => {
		const files = Array.from(event.target.files)
		if (files.length === 0) return

		const total_files = selectedFiles.value.length + files.length
		if (total_files > MAX_FILES) {
			alert(`一次最多只能上传${MAX_FILES}个文件`)
			event.target.value = ''
			return
		}

		// 立即将文件添加到列表，设置初始状态
		files.forEach(file => {
			selectedFiles.value.push({
				file: file,
				name: file.name,
				size: file.size,
				uploading: false,
				completed: false,
				progress: 0 // 单个文件的进度（可选）
			})
		})

		// 开始上传第一个文件
		if (selectedFiles.value.length > 0 && !isUploading.value) {
			await startFileUpload()
		}

		// 重置input，允许选择相同文件
		event.target.value = ''
	}

	const startFileUpload = async () => {
		if (selectedFiles.value.length === 0) {
			return
		}

		isUploading.value = true // 全局状态设为上传中（显示进度条）
		uploadProgress.value = 0

		const totalFiles = selectedFiles.value.length


		// 遍历所有文件，依次上传
		for (let i = 0; i < selectedFiles.value.length; i++) {
			const fileItem = selectedFiles.value[i]
			fileItem.uploading = true
			uploadingFileName.value = fileItem.name

			// 计算当前文件在整体进度中的权重
			const fileWeight = 100 / totalFiles
			const startProgress = (i / totalFiles) * 100

			// 上传当前文件
			await simulateFileUpload(i, startProgress, fileWeight)

			// 标记为完成
			fileItem.uploading = false
			fileItem.completed = true
			fileItem.progress = 100

			// 传递完整的文件信息而非仅文件名
			if (emit) {
				emit('file-uploaded', {
					name: fileItem.name,
					size: fileItem.size,
					file: fileItem.file, // 传递文件对象
					completed: fileItem.completed
				})
			}

			// 短暂延迟，让用户能看到每个文件的完成状态
			await delay(300)
		}

		// 所有文件上传完成后，延迟重置状态（让进度条显示完100%）
		setTimeout(() => {
			isUploading.value = false
			uploadProgress.value = 0
			uploadingFileName.value = ''
		}, 500)
	}

	// 模拟单个文件上传，支持整体进度
	const simulateFileUpload = (fileIndex, startProgress, fileWeight) => {
		return new Promise((resolve) => {
			let currentProgress = 0
			const step = 20 // 每个文件分5步完成 (100/5=20)

			// 清理之前的定时器
			if (progressTimer) {
				clearInterval(progressTimer)
				progressTimer = null
			}

			// 存储定时器ID
			progressTimer = setInterval(() => {
				currentProgress += step

				// 计算整体进度：起始进度 + 当前文件进度占整体进度的比例
				const overallProgress = startProgress + (currentProgress / 100) * fileWeight

				// 更新全局进度
				uploadProgress.value = Math.min(overallProgress, 100)

				// 更新单个文件进度
				selectedFiles.value[fileIndex].progress = currentProgress

				if (currentProgress >= 100) {
					clearInterval(progressTimer)
					progressTimer = null
					resolve() // 通知Promise完成
				}
			}, 500)
		})
	}

	// 拖拽上传支持
	const setupDragAndDrop = () => {
		if (!dropAreaRef || !dropAreaRef.value) return

		const dropArea = dropAreaRef.value

		const preventDefaults = (e) => {
			e.preventDefault()
			e.stopPropagation()
		}

		const highlight = () => {
			dropArea.style.borderColor = '#1890ff'
			dropArea.style.backgroundColor = '#f0f8ff'
		}

		const unhighlight = () => {
			dropArea.style.borderColor = '#d9d9d9'
			dropArea.style.backgroundColor = ''
		}

		const handleDrop = (e) => {
			preventDefaults(e)
			unhighlight()

			const files = e.dataTransfer.files
			if (files.length > 0) {
				// 模拟文件选择事件
				const event = {
					target: {
						files: files
					}
				}
				handleFileUpload(event)
			}
		}

		// 添加事件监听器
		;
		['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
			dropArea.addEventListener(eventName, preventDefaults, false)
		})

		;
		['dragenter', 'dragover'].forEach(eventName => {
			dropArea.addEventListener(eventName, highlight, false)
		})

		;
		['dragleave', 'drop'].forEach(eventName => {
			dropArea.addEventListener(eventName, unhighlight, false)
		})

		dropArea.addEventListener('drop', handleDrop, false)

		// 清理函数
		return () => {
			;
			['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
				dropArea.removeEventListener(eventName, preventDefaults, false)
			})

			;
			['dragenter', 'dragover'].forEach(eventName => {
				dropArea.removeEventListener(eventName, highlight, false)
			})

			;
			['dragleave', 'drop'].forEach(eventName => {
				dropArea.removeEventListener(eventName, unhighlight, false)
			})

			dropArea.removeEventListener('drop', handleDrop, false)
		}
	}

	// 在组件挂载时设置拖拽
	onMounted(() => {
		if (dropAreaRef) {
			const cleanup = setupDragAndDrop()
			if (cleanup) {
				onUnmounted(cleanup)
			}
		}
	})

	// 延迟函数
	const delay = (ms) => {
		return new Promise(resolve => setTimeout(resolve, ms))
	}

	const removeFile = (index) => {
		selectedFiles.value.splice(index, 1)
		// 如果没有文件了，重置上传状态
		if (selectedFiles.value.length === 0) {
			if (progressTimer) {
				clearInterval(progressTimer);
				progressTimer = null;
			}
			isUploading.value = false
			uploadProgress.value = 0
			uploadingFileName.value = ''
		}
	}

	// 清空方法中也清理定时器
	const clearFileList = () => {
		// 清理定时器，重置上传状态
		if (progressTimer) {
			clearInterval(progressTimer);
			progressTimer = null;
		}
		selectedFiles.value = []
		isUploading.value = false
		uploadProgress.value = 0
		uploadingFileName.value = ''
	}

	return {
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
		MAX_FILES // 导出最大文件数量
	}
}