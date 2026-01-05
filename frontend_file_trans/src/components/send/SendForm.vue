<!-- src/components/send/SendForm.vue -->
<template>
	<div class="send-form">
		<!-- 按钮组 -->
		<div class="button-group">
			<button class="send-but" :class="{ active: activeTab === 'file' }" @click="activeTab = 'file'">
				发送文件
			</button>
			<button class="pick-but" :class="{ active: activeTab === 'text' }" @click="activeTab = 'text'">
				发送文本
			</button>
		</div>

		<!-- 发送文件、文本区域 -->
		<div class="subsection">
			<!-- 发送文件区域 -->
			<div v-if="activeTab === 'file'">
				<FileUpload ref="fileUploadRef" @file-uploaded="handleFileUploaded" />
			</div>

			<!-- 发送文本区域 -->
			<div v-else class="text-area">
				<textarea v-model="textContent" placeholder="请输入要发送的文本内容..." class="text-input" rows="8"></textarea>
			</div>
		</div>

		<!-- 分隔线 -->
		<div class="divider"></div>

		<!-- 过期设置和取件选项 -->
		<div class="condition-section">
			<div class="condition">
				<ExpirySetting v-model:expiryDays="expiryDays" v-model:timeUnit="timeUnit" />

				<div class="sub-setion-cho">
					<div class="choose-note">
						<a href="javascript:void(0)" class="link" @click="goToReceive">需要取件？点击这里</a>
					</div>
				</div>
			</div>

			<!-- 提交按钮：添加加载状态和禁用 -->
			<div class="action-buttons">
				<button 
					class="btn-primary" 
					@click="handleSubmit"
					:disabled="isSubmitting"
				>
					<!-- 加载动画：isSubmitting时显示 -->
					<span class="spinner" v-if="isSubmitting"></span>
					<!-- 按钮文字：根据状态切换 -->
					<span v-if="!isSubmitting">安全发送</span>
					<span v-if="isSubmitting">发送中...</span>
				</button>
			</div>
		</div>

		<!-- 分隔线 -->
		<div class="divider"></div>

		<!-- 查看发送记录 -->
		<div class="record-section">
			<img src="@/assets/record_logo.png" alt="记录图标" class="record-icon">
			<button class="record-note" @click="$emit('show-records')">发件记录</button>
		</div>
	</div>
</template>

<script setup>
	import { ref } from 'vue'
	import { useRouter } from 'vue-router'
	import FileUpload from '@/components/common/FileUpload.vue'
	import ExpirySetting from '@/components/send/ExpirySetting.vue'
	import { formatFileSize } from '@/utils/fileUtils'

	const emit = defineEmits(['show-records', 'submit'])
	// 获取FileUpload组件的ref
	const fileUploadRef = ref(null)

	// 当前激活的标签页
	const activeTab = ref('file')
	// 文本内容
	const textContent = ref('')
	// 过期设置
	const expiryDays = ref(1)
	const timeUnit = ref('天')
	// 存储上传的文件信息
	const uploadedFiles = ref([])
	// 提交加载状态
	const isSubmitting = ref(false)

	// 接收完整的文件信息
	const handleFileUploaded = (fileInfo) => {
		console.log('文件上传完成:', fileInfo)
		// 检查是否已存在相同文件
		const exists = uploadedFiles.value.some(f => f.name === fileInfo.name && f.size === fileInfo.size)
		if (!exists) {
			uploadedFiles.value.push(fileInfo)
		}
	}

	// 页面跳转
	const router = useRouter()
	const goToReceive = () => {
		router.push('/receivePackage')
	}

	const handleSubmit = () => {
		if (activeTab.value === 'file') {
			// 文件发送验证逻辑
			if (uploadedFiles.value.length === 0) {
				alert('请先上传文件')
				return
			}
		} else if (activeTab.value === 'text' && !textContent.value.trim()) {
			alert('请输入要发送的文本内容')
			return
		}

		// 设置加载状态为true
		isSubmitting.value = true

		// 提交数据
		const submitData = {
			type: activeTab.value,
			text: activeTab.value === 'text' ? textContent.value : '',
			files: activeTab.value === 'file' ? uploadedFiles.value : [],
			expiryDays: expiryDays.value,
			timeUnit: timeUnit.value
		}

		emit('submit', submitData)
	}

	// 清空表单的方法（供父组件调用）
	const clearForm = () => {
		console.log('执行SendForm的clearForm方法')
		// 采用批量赋值，减少响应式更新的次数
		Promise.resolve().then(() => {
			textContent.value = ''
			uploadedFiles.value = []
			// 调用FileUpload的清空方法
			if (fileUploadRef.value && typeof fileUploadRef.value.clearFileList === 'function') {
				fileUploadRef.value.clearFileList()
			}
		});
	}

	// 重置提交状态（供父组件调用，恢复按钮初始状态）
	const resetSubmitting = () => {
		isSubmitting.value = false
	}

	// 暴露方法，让父组件可以调用
	defineExpose({
		clearForm,
		resetSubmitting // 新增暴露的重置方法
	})
</script>

<style scoped>
	.send-form {
		max-width: 500px;
		margin: 0 auto;
	}

	/* 按钮组样式 */
	.button-group {
		display: flex;
		justify-content: center;
		gap: 16px;
		margin-bottom: 24px;
	}

	.send-but,
	.pick-but {
		padding: 12px 24px;
		border: 2px solid #e0e0e0;
		border-radius: 8px;
		background-color: #f5f5f5;
		color: #666;
		font-size: 16px;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.3s ease;
	}

	.send-but:hover,
	.pick-but:hover {
		background-color: #e0e0e0;
		color: #333;
		border-color: #ccc;
	}

	.send-but.active,
	.pick-but.active {
		background-color: #1890ff;
		color: white;
		border-color: #1890ff;
	}

	.send-but.active:hover,
	.pick-but.active:hover {
		background-color: #40a9ff;
		border-color: #40a9ff;
	}

	.subsection {
		background: #fafafa;
		border-radius: 8px;
		padding: 0;
	}

	.text-area {
		width: 100%;
		box-sizing: border-box;
		padding: 20px;
	}

	.text-input {
		width: 100%;
		font-size: 14px;
		border: 2px dashed #d9d9d9;
		border-radius: 8px;
		padding: 16px;
		min-height: 160px;
		resize: vertical;
		background: transparent;
		transition: all 0.3s;
		box-sizing: border-box;
	}

	.text-input:focus {
		outline: none;
		border-color: #1890ff;
	}

	.divider {
		height: 1px;
		background: #f0f0f0;
		margin: 24px 0;
	}

	/* 条件设置区域 */
	.condition-section {
		width: 100%;
		margin-bottom: 24px;
	}

	.condition {
		display: flex;
		justify-content: space-between;
		gap: 20px;
		width: 100%;
	}

	.sub-setion-cho {
		flex: 1;
	}

	.choose-note {
		color: #666;
		margin-top: 60px;
		margin-left: 50px;
	}

	.link {
		color: #5500ff;
		text-decoration: none;
	}

	.link:hover {
		text-decoration: underline;
	}

	.action-buttons {
		text-align: center;
		margin-top: 32px;
	}

	/* 核心：提交按钮样式修改 */
	.btn-primary {
		background: #1890ff;
		color: white;
		border: none;
		border-radius: 6px;
		padding: 12px 32px;
		font-size: 16px;
		cursor: pointer;
		transition: background 0.3s;
		/* 新增：让按钮内元素居中对齐 */
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: 8px; /* 文字与加载动画的间距 */
	}

	/* 新增：鼠标悬停改为深色（深蓝光，可替换为其他深色） */
	.btn-primary:hover {
		/* 深蓝光（推荐）：#096dd9 */
		/* 深灰：#2c3e50 */
		/* 深绿：#0f5132 */
		background: #096dd9;
	}

	/* 新增：禁用状态样式（加载中时的按钮样式） */
	.btn-primary:disabled {
		background: #8c8c8c;
		cursor: not-allowed;
		opacity: 0.8;
	}

	/* 新增：加载动画样式（纯CSS旋转圆圈） */
	.spinner {
		width: 16px;
		height: 16px;
		border: 2px solid #fff;
		border-top: 2px solid transparent;
		border-radius: 50%;
		animation: spin 1s linear infinite;
	}

	/* 旋转动画关键帧 */
	@keyframes spin {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(360deg); }
	}

	.record-section {
		width: 100%;
		padding-right: 5px;
		box-sizing: border-box;
		text-align: right;
	}

	.record-container {
		display: inline-flex;
		align-items: center;
	}

	.record-icon {
		width: 20px;
		height: 20px;
		vertical-align: middle;
	}

	.record-note {
		margin-right: 1px;
		color: #999;
		cursor: pointer;
		transition: all 0.3s;
		background: transparent;
		border: none;
		font-size: 16px;
	}

	.record-note:hover {
		color: #5500ff;
	}

	/* 响应式设计 */
	@media (max-width: 768px) {
		.button-group {
			flex-direction: column;
			align-items: center;
		}

		.send-but,
		.pick-but {
			width: 100%;
			max-width: 200px;
		}

		.condition {
			flex-direction: column;
			gap: 10px;
		}

		.choose-note {
			margin-top: 20px;
			margin-left: 0;
			text-align: center;
		}
	}
</style>