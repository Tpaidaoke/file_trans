<template>
	<div v-if="visible" class="popup-mask">
		<div class="popup-content">
			<!-- 弹窗头部 -->
			<div class="popup-header">
				<h3>发件成功</h3>
				<button class="close-btn" @click="close">×</button>
			</div>

			<!-- 二维码区域 -->
			<div class="qr-section">
				<p class="section-title">取件二维码</p>
				<div class="qr-code">
					<img :src="fileInfo.qrCodeUrl" alt="取件二维码" />
				</div>
				<p class="qr-hint">扫描二维码快速取件</p>
			</div>

			<!-- 文件信息 -->
			<div class="info-section">
				<div class="info-row">
					<span class="info-label">文件名称：</span>
					<span class="info-value">{{ fileInfo.name }}</span>
				</div>

				<div class="info-row">
					<span class="info-label">文件大小：</span>
					<span class="info-value">{{ fileInfo.size }}</span>
				</div>

				<div class="info-row">
					<span class="info-label">过期时间：</span>
					<span class="info-value">{{ fileInfo.expireTip }} {{ fileInfo.expireUnit }}</span>
				</div>

				<div class="info-row">
					<span class="info-label">取件链接：</span>
					<span class="info-value link-text">{{ fileInfo.pickupLink }}</span>
				</div>

				<div class="info-row">
					<span class="info-label">取件日期：</span>
					<span class="info-value">{{ fileInfo.date }}</span>
				</div>
			</div>

			<!-- 操作按钮 -->
			<div class="action-buttons">
				<button class="btn-copy" @click="copyLink">复制链接</button>
				<button class="btn-close" @click="close">关闭</button>
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

	const close = () => {
		visible.value = false
	}

	// 复制取件链接到剪贴板
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
	}

	.popup-content {
		background: #fff;
		border-radius: 12px;
		width: 90%;
		max-width: 450px;
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
		overflow: hidden;
	}

	.popup-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 20px;
		background: #4f46e5;
		color: white;
	}

	.popup-header h3 {
		margin: 0;
		font-size: 18px;
		font-weight: 600;
	}

	.close-btn {
		background: none;
		border: none;
		color: white;
		font-size: 24px;
		cursor: pointer;
		line-height: 1;
		padding: 0;
		width: 30px;
		height: 30px;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: background-color 0.2s;
	}

	.close-btn:hover {
		background: rgba(255, 255, 255, 0.2);
	}

	.qr-section {
		padding: 25px;
		text-align: center;
		border-bottom: 1px solid #e5e7eb;
	}

	.section-title {
		margin: 0 0 15px 0;
		color: #374151;
		font-weight: 600;
		font-size: 16px;
	}

	.qr-code {
		width: 180px;
		height: 180px;
		margin: 0 auto;
		border: 1px solid #e5e7eb;
		padding: 10px;
		border-radius: 8px;
		background: white;
	}

	.qr-code img {
		width: 100%;
		height: 100%;
		object-fit: contain;
	}

	.qr-hint {
		margin: 10px 0 0 0;
		color: #6b7280;
		font-size: 14px;
	}

	.info-section {
		padding: 20px;
	}

	.info-row {
		display: flex;
		margin-bottom: 12px;
		align-items: flex-start;
	}

	.info-label {
		min-width: 80px;
		color: #6b7280;
		font-size: 14px;
		line-height: 1.5;
	}

	.info-value {
		flex: 1;
		color: #1f2937;
		font-size: 14px;
		line-height: 1.5;
		word-break: break-all;
	}

	.link-text {
		color: #4f46e5;
		font-weight: 500;
	}

	.action-buttons {
		display: flex;
		gap: 12px;
		padding: 20px;
		border-top: 1px solid #e5e7eb;
	}

	.btn-copy,
	.btn-close {
		flex: 1;
		padding: 12px;
		border-radius: 8px;
		font-size: 15px;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.2s;
		border: none;
	}

	.btn-copy {
		background: #4f46e5;
		color: white;
	}

	.btn-copy:hover {
		background: #4338ca;
		transform: translateY(-2px);
		box-shadow: 0 4px 12px rgba(79, 70, 229, 0.3);
	}

	.btn-close {
		background: #f3f4f6;
		color: #374151;
		border: 1px solid #d1d5db;
	}

	.btn-close:hover {
		background: #e5e7eb;
		transform: translateY(-2px);
	}

	/* 响应式调整 */
	@media (max-width: 480px) {
		.popup-content {
			width: 95%;
		}

		.qr-code {
			width: 150px;
			height: 150px;
		}

		.action-buttons {
			flex-direction: column;
		}
	}
</style>