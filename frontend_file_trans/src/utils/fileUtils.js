// src/utils/fileUtils.js

// 格式化文件大小
export function formatFileSize(bytes) {
	if (bytes === 0) return '0 Bytes'
	const k = 1024
	const sizes = ['Bytes', 'KB', 'MB', 'GB']
	const i = Math.floor(Math.log(bytes) / Math.log(k))
	return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// 验证文件类型
export function validateFileType(file, allowedTypes = []) {
	if (allowedTypes.length === 0) return true
	return allowedTypes.some(type => file.type.includes(type))
}

// 验证文件大小
export function validateFileSize(file, maxSizeInBytes) {
	return file.size <= maxSizeInBytes
}