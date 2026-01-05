<!-- src/views/sendPage.vue -->
<template>
	<div class="send-page">
		<!-- 页面头部 -->
		<CommonHeader title="发文件" />

		<!-- 主要内容区域 -->
		<div class="content">
			<SendForm ref="sendFormRef" @show-records="handleShowRecords" @submit="handleFormSubmit" />
		</div>

		<!-- 发件记录侧边栏 -->
		<RecordSidebar title="发件记录" empty-text="暂无发件记录" v-model="showRecordSidebar" :records="todaySendRecords" />

		<!-- 发件成功弹窗 -->
		<SendSuccessPopup v-model="showSendPopup" :file-info="sendFileInfo" />
	</div>
</template>

<script setup>
	import {
		ref,
		onMounted
	} from 'vue'
	import axios from 'axios'
	import {
		formatFileSize
	} from '@/utils/fileUtils'
	import {
		generateQrCodeUrl
	} from '@/utils/qrCodeUtils'
	import CommonHeader from '@/components/common/Header.vue'
	import SendForm from '@/components/send/SendForm.vue'
	import RecordSidebar from '@/components/common/RecordSidebar.vue'
	import SendSuccessPopup from '@/components/common/SendSuccessPopup.vue'

	// 控制弹窗显示,发件信息
	const showSendPopup = ref(false)
	const sendFileInfo = ref({})
	// 发件记录侧边栏显示状态
	const showRecordSidebar = ref(false)
	// 发件记录
	const todaySendRecords = ref([])
	// 获取SendForm组件的ref
	const sendFormRef = ref(null)

	// 生成6位随机数作为取件密钥
	const generateAccessKey = () => {
		const key = Math.floor(Math.random() * 1000000).toString().padStart(6, '0')
		return key
	}

	// 从后端获取发件记录
	const fetchSendRecords = async () => {
		try {
			const response = await axios.get('/api/v1/sendRecords')
			console.log(response.data)
			todaySendRecords.value = response.data.records || []
		} catch (error) {
			console.error('获取发件记录失败', error)
		}
	}

	// 处理点击发件记录的方法
	const handleShowRecords = async () => {
		// 等待获取最新的发件记录返回后显示
		await fetchSendRecords()
		console.log(todaySendRecords)
		showRecordSidebar.value = true
	}

	// 处理表单提交
	const handleFormSubmit = async (submitData) => {
		const accessKey = generateAccessKey()

		try {
			// 创建FormData对象，用于文件上传
			const formData = new FormData();

			formData.append("type", submitData.type);

			// 添加文本内容（如果是文本类型，后端需要额外处理）
			if (submitData.type === 'text') {
				formData.append('text', submitData.text);
				formData.append('textName', '文本文件.txt');
			}
			// 添加文件（如果是文件类型）
			else {
				// 后端期望的文件字段是"files"，直接追加文件对象
				submitData.files.forEach(fileItem => {
					formData.append('files', fileItem.file);
				});
			}

			formData.append("pickupCode", accessKey)
			// 拆分过期时间为数值和单位（适配后端参数）
			formData.append('expireTip', submitData.expiryDays);
			formData.append('expireUnit', submitData.timeUnit);

			// 调用后端API提交数据和文件（注意接口地址与后端路由匹配）
			const response = await axios.post('/api/v1/sendPackage', formData, {
				onUploadProgress: (progressEvent) => {
					const percentCompleted = Math.round(
						(progressEvent.loaded * 100) / (progressEvent.total || 1)
					);
				}
			});

			// 打印完整的后端返回数据
			console.log('后端返回完整响应:', response)
			console.log('后端返回数据:', response.data)
			console.log('后端返回状态码:', response.status)
			console.log('后端返回消息:', response.data.msg)
			console.log('后端返回文件下载URL:', response.data.fileDownloadURL)

			// 处理成功响应（适配后端返回结构）
			if (response.data.msg === '发送成功') {
				// 准备弹窗显示的文件信息
				let fileInfo = {};

				if (submitData.type === 'text') {
					// 文本类型处理
					fileInfo = {
						name: response.data.fileName,
						size: response.data.fileSize,
						date: new Date().toISOString().split('T')[0],
						expireTip: submitData.expiryDays,
						expireUnit: submitData.timeUnit,
						pickupLink: response.data.fileDownloadURL,
						pickupCode: accessKey,
						qrCodeUrl: ''
					};
				} else {
					// 处理文件类型
					fileInfo = {
						name: response.data.fileName,
						size: response.data.fileSize,
						date: new Date().toISOString().split('T')[0],
						expireTip: submitData.expiryDays,
						expireUnit: submitData.timeUnit,
						pickupLink: response.data.fileDownloadURL,
						pickupCode: accessKey,
						qrCodeUrl: ''
					};
				}

				// 更新文件信息并显示弹框
				sendFileInfo.value = fileInfo;
				showSendPopup.value = true;

				// 异步执行二维码生成
				Promise.resolve().then(async () => {
					const qrCodeText = fileInfo.pickupLink;
					const qrCodeDataUrl = await generateQrCodeUrl(qrCodeText);
					// 更新二维码地址（弹框中会自动刷新）
					sendFileInfo.value = {
						...sendFileInfo.value,
						qrCodeUrl: qrCodeDataUrl
					};
				});

				// 异步执行清空操作
				setTimeout(() => {
					if (sendFormRef.value && typeof sendFormRef.value.clearForm === 'function') {
						sendFormRef.value.clearForm();
					}
				}, 0);
			} else {
				alert('发送失败：' + (response.data.error || '未知错误'));
			}
		} catch (error) {
			// 更详细的错误信息
			if (error.response) {
				alert(`发送失败: ${error.response.data?.error || '服务器错误'}`);
			} else if (error.request) {
				alert('发送失败: 未收到服务器响应，请检查网络');
			} else {
				alert('发送失败: ' + error.message);
			}
		} finally {
			// 关键修复：无论成功还是失败，都重置提交状态
			// 使用 setTimeout 确保在下一个事件循环中执行，避免可能的时序问题
			setTimeout(() => {
				if (sendFormRef.value && typeof sendFormRef.value.resetSubmitting === 'function') {
					sendFormRef.value.resetSubmitting();
					console.log('已重置发送按钮状态');
				}
			}, 0);
		}
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