<template>
	<div v-if="visible" class="popup-mask">
		<div class="popup-content">
			<!-- 弹窗头部 -->
			<div class="popup-header">
				<h2>文件详情</h2>
				<button class="close-btn" @click="close">×</button>
			</div>

			<!-- 文件基本信息 -->
			<div class="file-basic-info">
				<div class="file-item">
					<span class="file-icon">📄</span>
					<div class="file-details">
						<span class="file-name">{{ fileInfo.name }}</span>
						<span class="file-meta">{{ fileInfo.size }} · {{ fileInfo.date }}</span>
					</div>
				</div>

				<div class="info-row">
					<span class="info-icon">⏰</span>
					<span class="info-text">{{ fileInfo.expireTip }} {{ fileInfo.expireUnit }}后过期</span>
				</div>
			</div>

			<!-- 分割线 -->
			<div class="divider"></div>

			<!-- 主要内容区域：取件码和二维码同行 -->
			<div class="main-content">
				<!-- 左侧：取件码和wget下载 -->
				<div class="left-section">
					<div class="pickup-code-section">
						<h3>取件码</h3>
						<div class="code-display">{{ props.fileInfo.pickupCode }}</div>
					</div>

					<div class="wget-section">
						<h3>wget下载</h3>
						<div class="wget-command" @click="copyWgetCommand">
							<code>{{ wgetCommand }}</code>
							<span class="copy-hint">点击复制wget命令</span>
						</div>
					</div>
				</div>

				<!-- 右侧：二维码 -->
				<div class="right-section">
					<div class="qr-code-area">
						<div class="qr-code">
							<img :src="fileInfo.qrCodeUrl" alt="取件二维码" />
						</div>
						<p class="qr-hint">扫描二维码快速取件</p>
					</div>
				</div>
			</div>

			<!-- 分割线 -->
			<div class="divider"></div>

			<!-- 底部：复制链接按钮 -->
			<div class="bottom-section">
				<button class="copy-link-btn" @click="copyLink">
					<span class="btn-icon">🔗</span>
					<span>复制取件链接</span>
				</button>
			</div>
		</div>
	</div>
</template>

<script setup>
	import {
		computed
	} from 'vue'

	const props = defineProps({
		modelValue: Boolean,
		fileInfo: {
			type: Object,
			default: () => ({})
		}
	})

	const emit = defineEmits(['update:modelValue'])

	const visible = computed({
		get: () => props.modelValue,
		set: (val) => emit('update:modelValue', val)
	})

	// 生成wget命令
	const wgetCommand = computed(() => {
		if (props.fileInfo.pickupLink) {
			return `wget "${props.fileInfo.pickupLink}"`
		}
		return 'wget "https://example.com/download/file123"'
	})

	const close = () => {
		visible.value = false
	}

	// 复制取件链接
	const copyLink = () => {
		if (!props.fileInfo.pickupLink) return

		navigator.clipboard.writeText(props.fileInfo.pickupLink)
			.then(() => {
				alert('取件链接已复制到剪贴板！')
			})
			.catch(err => {
				console.error('复制失败:', err)
				// 备用方法
				const textArea = document.createElement('textarea')
				textArea.value = props.fileInfo.pickupLink
				document.body.appendChild(textArea)
				textArea.select()
				document.execCommand('copy')
				document.body.removeChild(textArea)
				alert('取件链接已复制到剪贴板！')
			})
	}

	// 复制wget命令
	const copyWgetCommand = () => {
		navigator.clipboard.writeText(wgetCommand.value)
			.then(() => {
				alert('wget命令已复制到剪贴板！')
			})
			.catch(err => {
				console.error('复制失败:', err)
				// 备用方法
				const textArea = document.createElement('textarea')
				textArea.value = wgetCommand.value
				document.body.appendChild(textArea)
				textArea.select()
				document.execCommand('copy')
				document.body.removeChild(textArea)
				alert('wget命令已复制到剪贴板！')
			})
	}
</script>

<style scoped>
	.popup-mask {
		position: fixed;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background: rgba(0, 0, 0, 0.5);
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 1000;
		overflow-y: auto;
		padding: 20px;
	}

	.popup-content {
		background: white;
		border-radius: 16px;
		width: 100%;
		max-width: 700px;
		/* 增加宽度 */
		max-height: 90vh;
		/* 限制最大高度 */
		overflow-y: auto;
		/* 内部滚动 */
		box-shadow: 0 10px 40px rgba(0, 0, 0, 0.2);
		display: flex;
		flex-direction: column;
	}

	.popup-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 20px 24px;
		border-bottom: 1px solid #f0f0f0;
		flex-shrink: 0;
	}

	.popup-header h2 {
		margin: 0;
		font-size: 20px;
		font-weight: 600;
		color: #333;
	}

	.close-btn {
		background: none;
		border: none;
		font-size: 28px;
		color: #666;
		cursor: pointer;
		width: 30px;
		height: 30px;
		display: flex;
		align-items: center;
		justify-content: center;
		border-radius: 50%;
		transition: background-color 0.2s;
		flex-shrink: 0;
	}

	.close-btn:hover {
		background: #f5f5f5;
	}

	.file-basic-info {
		padding: 20px 24px;
		flex-shrink: 0;
	}

	.file-item {
		display: flex;
		align-items: center;
		gap: 12px;
		margin-bottom: 16px;
	}

	.file-icon {
		font-size: 32px;
		width: 50px;
		height: 50px;
		background: #f0f7ff;
		border-radius: 10px;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-shrink: 0;
	}

	.file-details {
		display: flex;
		flex-direction: column;
		min-width: 0;
	}

	.file-name {
		font-weight: 500;
		color: #333;
		font-size: 16px;
		margin-bottom: 4px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	.file-meta {
		color: #666;
		font-size: 14px;
	}

	.info-row {
		display: flex;
		align-items: center;
		gap: 10px;
		color: #666;
		font-size: 14px;
	}

	.info-icon {
		width: 20px;
		text-align: center;
		flex-shrink: 0;
	}

	.divider {
		height: 1px;
		background: #f0f0f0;
		margin: 0 24px;
		flex-shrink: 0;
	}

	.main-content {
		padding: 20px 24px;
		display: flex;
		gap: 24px;
		align-items: flex-start;
		flex: 1;
		min-height: 0;
	}

	.left-section {
		flex: 1;
		min-width: 0;
	}

	.pickup-code-section {
		margin-bottom: 20px;
	}

	.pickup-code-section h3 {
		margin: 0 0 12px 0;
		font-size: 16px;
		color: #333;
		font-weight: 600;
	}

	.code-display {
		background: #f8f9fa;
		border: 2px solid #e9ecef;
		border-radius: 8px;
		padding: 12px;
		font-size: 24px;
		font-weight: 700;
		text-align: center;
		letter-spacing: 2px;
		color: #333;
		word-break: break-all;
	}

	.wget-section {
		margin-top: 16px;
	}

	.wget-section h4 {
		margin: 0 0 8px 0;
		font-size: 14px;
		color: #666;
		font-weight: 500;
	}

	.wget-command {
		background: #f8f9fa;
		border: 1px solid #e9ecef;
		border-radius: 8px;
		padding: 12px;
		cursor: pointer;
		transition: all 0.2s;
	}

	.wget-command:hover {
		background: #f1f3f4;
		border-color: #4dabf7;
	}

	.wget-command code {
		font-family: 'Courier New', monospace;
		font-size: 13px;
		color: #333;
		word-break: break-all;
		display: block;
		margin-bottom: 4px;
	}

	.copy-hint {
		font-size: 12px;
		color: #4dabf7;
		text-align: right;
		display: block;
	}

	.right-section {
		display: flex;
		flex-direction: column;
		align-items: center;
		min-width: 160px;
		flex-shrink: 0;
	}

	.qr-code-area {
		text-align: center;
	}

	.qr-code {
		width: 140px;
		height: 140px;
		border: 1px solid #f0f0f0;
		padding: 8px;
		border-radius: 8px;
		background: white;
		margin-bottom: 8px;
	}

	.qr-code img {
		width: 100%;
		height: 100%;
		object-fit: contain;
	}

	.qr-hint {
		margin: 0;
		color: #666;
		font-size: 13px;
		white-space: nowrap;
	}

	.bottom-section {
		padding: 20px 24px;
		border-top: 1px solid #f0f0f0;
		flex-shrink: 0;
	}

	.copy-link-btn {
		width: 100%;
		background: #4dabf7;
		border: none;
		border-radius: 10px;
		padding: 14px;
		color: white;
		font-size: 16px;
		font-weight: 500;
		cursor: pointer;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 8px;
		transition: all 0.2s;
	}

	.copy-link-btn:hover {
		background: #339af0;
		transform: translateY(-2px);
		box-shadow: 0 4px 12px rgba(77, 171, 247, 0.3);
	}

	.btn-icon {
		font-size: 18px;
	}

	/* 响应式调整 */
	@media (max-width: 480px) {
		.popup-content {
			max-width: 95%;
		}

		.popup-header,
		.file-basic-info,
		.main-content,
		.bottom-section {
			padding: 16px;
		}

		.main-content {
			flex-direction: column;
			gap: 20px;
		}

		.right-section {
			min-width: 100%;
			order: -1;
			/* 在移动端将二维码放在顶部 */
		}

		.code-display {
			font-size: 20px;
			padding: 10px;
		}

		.qr-code {
			width: 120px;
			height: 120px;
		}
	}
</style>