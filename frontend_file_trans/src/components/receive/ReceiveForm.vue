<template>
	<div class="rec-form">
		<h3 class="form-title">取件码</h3>
		<div class="input-group">
			<textarea 
				v-model="pickupCode" 
				placeholder="请输入6位取件码..." 
				class="text-input" 
				rows="1"
				@keyup.enter="handlePickup" 
				@input="handleInput" 
				:maxlength="6"
				:disabled="isLoading"
			/>
			<div class="char-count" :class="{ 'max-reached': isMaxLength }">
				{{ pickupCode.length }}/6
			</div>
		</div>
		<button 
			class="rec-btn" 
			@click="handlePickup" 
			:disabled="!isValidPickupCode || isLoading"
		>
			<template v-if="isLoading">取件中...</template>
			<template v-else>取件</template>
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

	<!-- 错误提示弹窗 -->
	<div v-if="showError" class="error-toast">
		{{ errorMessage }}
	</div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const emit = defineEmits(['show-records', 'pickup-success', 'submit'])

// 页面路由跳转
const router = useRouter()
const goToSend = () => {
  router.push('/sendPackage')
}

// 取件码状态管理
const pickupCode = ref('')
const isLoading = ref(false)  // 加载状态
const showError = ref(false)  // 错误提示显示状态
const errorMessage = ref('')  // 错误提示信息

// 计算属性：检查是否为有效的6位数字取件码
const isValidPickupCode = computed(() => {
	return pickupCode.value.length === 6 && /^\d+$/.test(pickupCode.value)
})

// 计算属性：是否达到最大长度
const isMaxLength = computed(() => {
	return pickupCode.value.length === 6
})

// 输入处理：只允许数字且限制长度
const handleInput = (event) => {
	// 只允许输入数字
	let value = event.target.value.replace(/\D/g, '')

	// 限制为6位数字（修正注释错误）
	if (value.length > 6) {
		value = value.slice(0, 6)
	}

	pickupCode.value = value
}

// 隐藏错误提示
const hideError = () => {
	showError.value = false
	errorMessage.value = ''
}

// 处理取件逻辑
const handlePickup = async () => {
	// 基础校验
	if (!isValidPickupCode.value) {
		alert('请输入6位数字取件码')
		return
	}

	// 清空之前的错误提示
	hideError()
	
	try {
		// 开始加载
		isLoading.value = true
		
		// 调用后端接口（GET请求，传递取件码参数）
		const response = await axios.get('/api/v1/receivePackage', {
			params: {
				pickupCode: pickupCode.value  // 匹配后端接收的查询参数名
			},
			timeout: 10000  // 10秒超时设置
		})
		console.log(response.data)
		// 接口调用成功
		if (response.data && response.data.fileDownloadUrl) {
			// 清空输入框
			pickupCode.value = ''
			// 通知父组件取件成功，传递文件信息
			emit('pickup-success', response.data)
		} else {
			throw new Error('获取文件信息失败')
		}

	} catch (error) {
		// 错误处理
		if (error.response) {
			// 后端返回错误状态码
			switch (error.response.status) {
				case 400:
					errorMessage.value = '请输入正确的取件码'
					break
				case 404:
					errorMessage.value = '取件码已过期或不存在'
					break
				case 500:
					errorMessage.value = '服务器繁忙，请稍后再试'
					break
				default:
					errorMessage.value = '取件失败，请重试'
			}
		} else if (error.request) {
			// 网络错误
			errorMessage.value = '网络连接失败，请检查网络'
		} else {
			// 其他错误
			errorMessage.value = error.message || '取件失败，请重试'
		}
		// 显示错误提示
		showError.value = true
		// 3秒后自动隐藏错误提示
		setTimeout(hideError, 3000)
	} finally {
		// 结束加载
		isLoading.value = false
	}
}
</script>

<style scoped>
	/* 原有样式保持不变 */
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
		font-size: 16px;
		letter-spacing: 2px;
		text-align: center;
	}

	.text-input:focus {
		outline: none;
		border-color: #1890ff;
		box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
	}

	.text-input::placeholder {
		color: #999;
		letter-spacing: normal;
		text-align: center;
	}

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

	/* 新增错误提示样式 */
	.error-toast {
		position: fixed;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		background: rgba(0, 0, 0, 0.7);
		color: white;
		padding: 12px 20px;
		border-radius: 8px;
		font-size: 14px;
		z-index: 1000;
		animation: fadeInOut 3s ease-in-out;
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

	/* 动画效果 */
	@keyframes fadeInOut {
		0% { opacity: 0; }
		20% { opacity: 1; }
		80% { opacity: 1; }
		100% { opacity: 0; }
	}
</style>