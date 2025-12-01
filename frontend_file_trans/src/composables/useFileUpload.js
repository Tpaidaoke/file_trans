// src/composables/useFileUpload.js
import {
	ref
} from 'vue'
import {
	useSendRecords
} from './useSendRecords'

export function useFileUpload(emit) {
	const fileInput = ref(null)
	const selectedFiles = ref([])
	const isUploading = ref(false)
	const uploadProgress = ref(0)
	const uploadingFileName = ref('')

	const {
		addRecord
	} = useSendRecords()

	const triggerFileInput = () => {
		fileInput.value.click()
	}

	const handleFileUpload = async (event) => {
		const files = Array.from(event.target.files)
		if (files.length === 0) return

		// 立即将文件添加到列表，设置初始状态
		files.forEach(file => {
			selectedFiles.value.push({
				file: file,
				name: file.name,
				size: file.size,
				uploading: false,
				completed: false,
				progress: 0
			})
		})

		// 开始上传第一个文件
		if (selectedFiles.value.length > 0) {
			await startFileUpload()
		}

		// 重置input，允许选择相同文件
		event.target.value = ''
	}

	const startFileUpload = async () => {
		// 找到第一个未完成上传的文件
		const pendingFileIndex = selectedFiles.value.findIndex(item => !item.completed && !item.uploading)

		if (pendingFileIndex === -1) {
			isUploading.value = false
			return
		}

		isUploading.value = true
		const fileItem = selectedFiles.value[pendingFileIndex]
		fileItem.uploading = true
		uploadingFileName.value = fileItem.name

		// 模拟上传进度
		await simulateUploadProgress(pendingFileIndex)

		// 标记为完成
		fileItem.uploading = false
		fileItem.completed = true
		fileItem.progress = 100

		// 添加到发件记录
		addRecord(fileItem.name)

		// 通知父组件文件上传完成
		if (emit) {
			emit('file-uploaded', fileItem.name)
		}

		// 继续上传下一个文件
		setTimeout(() => {
			startFileUpload()
		}, 500)
	}

	const simulateUploadProgress = (fileIndex) => {
		return new Promise((resolve) => {
			const fileItem = selectedFiles.value[fileIndex]
			fileItem.progress = 0
			uploadProgress.value = 0

			const interval = setInterval(() => {
				fileItem.progress += Math.floor(Math.random() * 10) + 1
				uploadProgress.value = fileItem.progress

				if (fileItem.progress >= 100) {
					fileItem.progress = 100
					uploadProgress.value = 100
					clearInterval(interval)
					setTimeout(() => {
						resolve()
					}, 500)
				}
			}, 300)
		})
	}

	const removeFile = (index) => {
		selectedFiles.value.splice(index, 1)
		// 如果没有文件了，重置上传状态
		if (selectedFiles.value.length === 0) {
			isUploading.value = false
			uploadProgress.value = 0
		}
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
		simulateUploadProgress
	}
}