<template>
	<div class="rec-form">
		<h3 class="form-title">取件码</h3>
		<div class="input-group">
			<textarea v-model="pickupCode" placeholder="请输入5位取件码..." class="text-input" rows="1"
				@keyup.enter="handlePickup" @input="handleInput" :maxlength="5" />
			<div class="char-count" :class="{ 'max-reached': isMaxLength }">
				{{ pickupCode.length }}/5
			</div>
		</div>
		<button class="rec-btn" @click="handlePickup" :disabled="!isValidPickupCode">
			取件
		</button>
	</div>
	<!-- 分隔线 -->
	<div class="divider"></div>

	<!-- 查看取件记录 -->
	<div class="record-section">
		<div class="choose-note">
			<a href="javascript:void(0)" class="link" @click="goToSend">需要发件？点击这里</a>
		</div>
		<div class="record-note">
			<img src="@/assets/record_logo.png" alt="记录图标" class="record-icon">
			<button class="record-btn" @click="$emit('show-records')">取件记录</button>
		</div>
	</div>
</template>

<script setup>
	import { ref,computed } from 'vue'
	import { useRouter } from 'vue-router'
	import { usePickupRecords } from '@/composables/usePickupRecords'

	const {
		addPickupRecord
	} = usePickupRecords()
	const emit = defineEmits(['show-records', 'pickup-success'])
	
	// 页面跳转
	const router = useRouter()
	
	const goToSend = () => {
	  router.push('/send')
	}

	const pickupCode = ref('')

	// 计算属性：检查是否为有效的5位数字取件码
	const isValidPickupCode = computed(() => {
		return pickupCode.value.length === 5 && /^\d+$/.test(pickupCode.value)
	})

	// 计算属性：是否达到最大长度
	const isMaxLength = computed(() => {
		return pickupCode.value.length === 5
	})

	const handleInput = (event) => {
		// 只允许输入数字
		let value = event.target.value.replace(/\D/g, '')

		// 限制为5位数字
		if (value.length > 5) {
			value = value.slice(0, 5)
		}

		pickupCode.value = value
	}

	const handlePickup = () => {
		if (!isValidPickupCode.value) {
			alert('请输入5位数字取件码')
			return
		}

		try {
			// 模拟API调用验证取件码
			console.log('验证取件码:', pickupCode.value)

			// 这里调用实际的后端接口
			// const response = await api.validatePickupCode(pickupCode.value)

			// 模拟成功响应
			const fileName = `文件_${pickupCode.value}.txt` // 实际应该从接口返回

			// 添加到取件记录
			addPickupRecord(fileName)

			// 通知父组件取件成功
			emit('pickup-success', fileName)

			alert(`取件成功: ${fileName}`)

			// 清空输入框
			pickupCode.value = ''

		} catch (error) {
			alert('取件码无效或文件已过期')
		}
	}
</script>

<style scoped>
	.rec-form {
		max-width: 500px;
		margin: 0 auto;
		padding: 20px;
	}

	.form-title {
		font-size: 18px;
		font-weight: 600;
		margin-bottom: 16px;
		color: #1a1a1a;
		text-align: left;
	}

	.input-group {
		position: relative;
		display: flex;
		gap: 12px;
		align-items: flex-start;
	}

	.text-input {
		flex: 1;
		font-size: 14px;
		border: 2px solid #d9d9d9;
		border-radius: 8px;
		padding: 12px 16px;
		resize: none;
		min-height: 48px;
		transition: all 0.3s;
		box-sizing: border-box;
		font-family: 'Courier New', monospace;
		/* 等宽字体，数字显示更整齐 */
		font-size: 16px;
		letter-spacing: 2px;
		/* 数字间距，更易读 */
		text-align: center;
		/* 居中显示 */
	}

	.text-input:focus {
		outline: none;
		border-color: #1890ff;
		box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
	}

	.text-input::placeholder {
		color: #999;
		letter-spacing: normal;
		/* placeholder 不需要字母间距 */
		text-align: center;
	}

	/* 字符计数样式 */
	.char-count {
		position: absolute;
		right: 12px;
		top: 12px;
		font-size: 12px;
		color: #999;
		background: rgba(255, 255, 255, 0.9);
		padding: 2px 6px;
		border-radius: 4px;
		transition: all 0.3s;
	}

	.char-count.max-reached {
		color: #52c41a;
		font-weight: 500;
	}

	.rec-btn {
		width: 100%;
		margin-top: 15px;
		padding: 12px 24px;
		border: 2px solid #1890ff;
		border-radius: 8px;
		background-color: #1890ff;
		color: white;
		font-size: 16px;
		font-weight: 500;
		cursor: pointer;
		transition: all 0.3s ease;
		min-width: 80px;
		height: 48px;
		white-space: nowrap;
	}

	.rec-btn:not(:disabled):hover {
		background-color: #40a9ff;
		border-color: #40a9ff;
		transform: translateY(-1px);
		box-shadow: 0 4px 8px rgba(24, 144, 255, 0.3);
	}

	.rec-btn:not(:disabled):active {
		transform: translateY(0);
		box-shadow: 0 2px 4px rgba(24, 144, 255, 0.3);
	}

	.rec-btn:disabled {
		background-color: #f5f5f5;
		border-color: #d9d9d9;
		color: #ccc;
		cursor: not-allowed;
		transform: none;
		box-shadow: none;
	}

	/* 当按钮可用时的特殊样式 */
	.rec-btn:not(:disabled) {
		background-color: #52c41a;
		border-color: #52c41a;
	}

	.rec-btn:not(:disabled):hover {
		background-color: #73d13d;
		border-color: #73d13d;
		box-shadow: 0 4px 8px rgba(82, 196, 26, 0.3);
	}

	.divider {
		height: 1px;
		background: #f0f0f0;
		margin: 24px 0;
	}

	.record-section {
		display: flex;
		justify-content: space-between;
		gap: 20px;
	}
	
	.link {
		color: #5500ff;
		text-decoration: none;
	}
	
	.link:hover {
		text-decoration: underline;
	}

	.record-note {
		padding-right: 5px;
		box-sizing: border-box;
		text-align: right;
	}

	.record-icon {
		width: 20px;
		height: 20px;
		vertical-align: middle;
	}

	.record-btn {
		margin-right: 1px;
		color: #999;
		cursor: pointer;
		transition: all 0.3s;
		background: transparent;
		border: none;
		font-size: 16px;
	}

	.record-btn:hover {
		color: #5500ff;
	}

	/* 响应式设计 */
	@media (max-width: 768px) {
		.rec-form {
			padding: 16px;
		}

		.input-group {
			flex-direction: column;
		}

		.rec-btn {
			width: 100%;
			height: 44px;
		}

		.record-btn {
			margin-top: 20px;
			margin-left: 0;
			text-align: center;
		}

		.char-count {
			position: relative;
			right: auto;
			top: auto;
			align-self: flex-end;
			margin-top: 8px;
		}
	}
</style>