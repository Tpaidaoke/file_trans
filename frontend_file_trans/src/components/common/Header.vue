<template>
	<div class="header">
		<div class="logo-container" @click="handleLogoClick">
			<img class="logo" :src="currentLogo" alt="logo" :key="logoKey">
			<div class="progress-ring" :style="progressStyle"></div>
			<!-- <div class="time-indicator">{{ timeLeft }}s</div> -->
		</div>
		<div class="title">{{ title }}</div>
	</div>
</template>

<script setup>
	import {
		ref,
		computed,
		onMounted,
		onUnmounted
	} from 'vue'

	const props = defineProps({
		title: {
			type: String,
			default: '发取文件'
		}
	})

	const logos = [
		"src/assets/logo.png",
		"src/assets/view_logo.png",
		"src/assets/send_logo.png",
		"src/assets/pickup_logo.png",
		"src/assets/01.png",
		"src/assets/02.png",
		"src/assets/03.png",
		"src/assets/04.png",
		"src/assets/05.png",
		"src/assets/06.png",
		"src/assets/07.png",
		"src/assets/08.png",
		"src/assets/09.png"
	]

	const currentLogo = ref(logos[0])
	const logoKey = ref(0)
	const timeLeft = ref(15) // 剩余时间
	let autoChangeInterval = null
	let countdownInterval = null
	let isClicking = false

	// 计算进度环的样式
	const progressStyle = computed(() => {
		const circumference = 2 * Math.PI * 22 // 半径22
		const offset = circumference * (1 - timeLeft.value / 15)
		return {
			strokeDasharray: `${circumference} ${circumference}`,
			strokeDashoffset: offset
		}
	})

	// 随机变换logo并带有动画效果
	const randomLogoWithEffect = () => {
		if (isClicking) return

		isClicking = true

		let count = 0
		const maxCount = 15
		const shuffleInterval = setInterval(() => {
			const randomIndex = Math.floor(Math.random() * logos.length)
			currentLogo.value = logos[randomIndex]
			count++

			if (count >= maxCount) {
				clearInterval(shuffleInterval)
				const finalIndex = Math.floor(Math.random() * logos.length)
				currentLogo.value = logos[finalIndex]
				logoKey.value++

				setTimeout(() => {
					isClicking = false
				}, 500)
			}
		}, 50)
	}

	// 重置倒计时
	const resetCountdown = () => {
		timeLeft.value = 15
	}

	// 开始倒计时
	const startCountdown = () => {
		if (countdownInterval) {
			clearInterval(countdownInterval)
		}

		countdownInterval = setInterval(() => {
			timeLeft.value--

			if (timeLeft.value <= 0) {
				clearInterval(countdownInterval)
			}
		}, 1000)
	}

	// 处理点击事件
	const handleLogoClick = () => {
		// 清除自动切换定时器
		if (autoChangeInterval) {
			clearInterval(autoChangeInterval)
		}

		// 重置倒计时
		resetCountdown()

		// 执行变换动画
		randomLogoWithEffect()

		// 重新启动自动切换
		startAutoChange()
	}

	// 启动自动切换
	const startAutoChange = () => {
		if (autoChangeInterval) {
			clearInterval(autoChangeInterval)
		}

		// 重置并开始倒计时
		resetCountdown()
		startCountdown()

		// 设置15秒自动切换
		autoChangeInterval = setInterval(() => {
			randomLogoWithEffect()
			resetCountdown()
			startCountdown()
		}, 15000)
	}

	// 组件挂载时启动
	onMounted(() => {
		// 延迟3秒后开始自动切换
		setTimeout(() => {
			startAutoChange()
		}, 3000)
	})

	// 组件卸载时清理
	onUnmounted(() => {
		if (autoChangeInterval) clearInterval(autoChangeInterval)
		if (countdownInterval) clearInterval(countdownInterval)
	})
</script>

<style scoped>
	.header {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-top: 20px;
		position: relative;
	}

	.logo-container {
		margin-right: 12px;
		cursor: pointer;
		position: relative;
		width: 56px;
		height: 56px;
	}

	.logo {
		width: 48px;
		height: 48px;
		transition: all 0.3s ease;
		border-radius: 8px;
		position: absolute;
		top: 4px;
		left: 4px;
		z-index: 2;
	}

	.logo:hover {
		transform: rotate(15deg) scale(1.1);
		box-shadow: 0 0 15px rgba(75, 116, 227, 0.5);
	}

	/* 进度环 */
	.progress-ring {
		position: absolute;
		top: 0;
		left: 0;
		width: 56px;
		height: 56px;
		z-index: 1;
	}

	.progress-ring::before {
		content: '';
		position: absolute;
		width: 100%;
		height: 100%;
		border-radius: 50%;
		border: 2px solid #e0e0e0;
	}

	.progress-ring::after {
		content: '';
		position: absolute;
		width: 100%;
		height: 100%;
		border-radius: 50%;
		border: 2px solid #4b74e3;
		border-top-color: transparent;
		border-right-color: transparent;
		transform: rotate(-90deg);
		transition: stroke-dashoffset 1s linear;
		animation: rotate 15s linear infinite;
	}

	/* 时间指示器 */
	/* .time-indicator {
		position: absolute;
		bottom: -20px;
		left: 50%;
		transform: translateX(-50%);
		font-size: 12px;
		color: #666;
		font-weight: bold;
		background: rgba(255, 255, 255, 0.9);
		padding: 2px 6px;
		border-radius: 10px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
		z-index: 3;
		min-width: 30px;
		text-align: center;
	} */

	/* Vue的transition动画 */
	.logo {
		animation: logoAppear 0.5s ease;
	}

	@keyframes logoAppear {
		0% {
			opacity: 0;
			transform: scale(0.8) rotate(-180deg);
		}

		70% {
			transform: scale(1.1) rotate(10deg);
		}

		100% {
			opacity: 1;
			transform: scale(1) rotate(0deg);
		}
	}

	@keyframes rotate {
		from {
			transform: rotate(-90deg);
		}

		to {
			transform: rotate(270deg);
		}
	}

	.title {
		font-size: 40px;
		font-weight: 600;
		color: #4b74e3;
		line-height: 48px;
	}
</style>