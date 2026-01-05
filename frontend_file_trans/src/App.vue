<script setup>
	import {
		ref,
		onMounted
	} from 'vue';

	// 状态管理
	const isDark = ref(true);

	// 切换主题函数
	const toggleTheme = () => {
		isDark.value = !isDark.value;
		updateTheme();
	};

	// 更新 DOM 和 LocalStorage
	const updateTheme = () => {
		const themeValue = isDark.value ? 'dark' : 'light';
		// 设置 html 标签的 data-theme 属性，方便 CSS 选择器定位
		document.documentElement.setAttribute('data-theme', themeValue);
		localStorage.setItem('theme', themeValue);
	};

	// 初始化：检查本地存储或系统偏好
	onMounted(() => {
		const savedTheme = localStorage.getItem('theme');
		if (savedTheme) {
			isDark.value = savedTheme === 'dark';
		} else {
			// 如果没有存过，检查系统偏好
			isDark.value = window.matchMedia('(prefers-color-scheme: dark)').matches;
		}
		updateTheme();
	});
</script>

<template>
	<div id="app">
		<!-- 主题切换按钮 -->
		<div class="theme-switch-wrapper" @click="toggleTheme">
			<div class="theme-switch" :class="{ 'active': !isDark }">
				<div class="switch-handle">
					<!-- 太阳图标 (浅色模式显示) -->
					<svg v-if="!isDark" class="icon sun" viewBox="0 0 24 24" fill="none"
						xmlns="http://www.w3.org/2000/svg">
						<circle cx="12" cy="12" r="5" stroke="currentColor" stroke-width="2" />
						<path
							d="M12 2V4M12 20V22M4.92999 4.92999L6.34 6.34M17.66 17.66L19.07 19.07M2 12H4M20 12H22M4.92999 19.07L6.34 17.66M17.66 6.34L19.07 4.92999"
							stroke="currentColor" stroke-width="2" stroke-linecap="round" />
					</svg>
					<!-- 月亮图标 (深色模式显示) -->
					<svg v-else class="icon moon" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
						<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79Z" stroke="currentColor" stroke-width="2"
							stroke-linecap="round" stroke-linejoin="round" />
					</svg>
				</div>
			</div>
		</div>

		<router-view />
	</div>
</template>

<style>
	:root {
		/* --- 深色模式 (默认/原有) --- */
		--bg-primary: #0a0a0f;
		--bg-secondary: #1a1a2e;
		--bg-accent: #16213e;
		--bg-highlight: #0f3460;
		--text-color: #ffffff;
		--tech-blue: #00aaff;
		--tech-purple: #a100ff;
		--grid-color: rgba(0, 200, 255, 0.1);
		--card-bg: rgba(16, 22, 36, 0.8);
		--card-border: rgba(0, 200, 255, 0.2);
		--shadow-color: rgba(0, 0, 0, 0.3);
	}

	/* --- 浅色模式重写变量 --- */
	:root[data-theme="light"] {
		--bg-primary: #f0f2f5;
		/* 浅灰背景 */
		--bg-secondary: #ffffff;
		/* 纯白 */
		--bg-accent: #e6f7ff;
		/* 淡蓝 */
		--bg-highlight: #bae7ff;
		/* 亮蓝 */
		--text-color: #1a1a2e;
		/* 深色文字 */
		--tech-blue: #0077cc;
		/* 稍微加深的蓝色以适应浅色背景 */
		--tech-purple: #8000cc;
		--grid-color: rgba(0, 100, 200, 0.08);
		/* 网格颜色变深一点点 */
		--card-bg: rgba(255, 255, 255, 0.7);
		--card-border: rgba(0, 170, 255, 0.2);
		--shadow-color: rgba(0, 100, 200, 0.15);
	}

	* {
		margin: 0;
		padding: 0;
		box-sizing: border-box;
	}

	#app {
		min-height: 100vh;
		font-family: Avenir, Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;

		/* 1. 将硬编码的 #fff 改为变量 */
		color: var(--text-color);

		/* 增加颜色过渡效果，使切换更丝滑 */
		transition: color 0.5s ease, background 0.5s ease;

		/* 使用CSS变量 */
		background: linear-gradient(135deg,
				var(--bg-primary) 0%,
				var(--bg-secondary) 30%,
				var(--bg-accent) 70%,
				var(--bg-highlight) 100%);

		/* 动态网格背景 */
		/* 注意：如果不重置background-image，上面的linear-gradient可能会被覆盖，
       这里使用多重背景或保持原样逻辑 */
		background-image:
			linear-gradient(var(--grid-color) 1px, transparent 1px),
			linear-gradient(90deg, var(--grid-color) 1px, transparent 1px),
			linear-gradient(135deg,
				var(--bg-primary) 0%,
				var(--bg-secondary) 30%,
				var(--bg-accent) 70%,
				var(--bg-highlight) 100%);

		background-size: 50px 50px, 50px 50px, 100% 100%;
		/* 修正动画：只移动网格，不移动背景色 */
		background-attachment: scroll, scroll, fixed;
		animation: gridMove 30s linear infinite;

		position: relative;
		overflow-x: hidden;
	}

	/* 数据流效果 */
	#app::after {
		content: '';
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		/* 调整数据流颜色，使其在浅色模式下也能稍微看见一点 */
		background: linear-gradient(90deg,
				transparent 0%,
				rgba(0, 200, 255, 0.05) 50%,
				transparent 100%);
		animation: dataFlow 10s linear infinite;
		pointer-events: none;
		z-index: 1;
	}

	@keyframes gridMove {
		0% {
			background-position: 0 0, 0 0, 0 0;
		}

		100% {
			background-position: 50px 50px, 50px 50px, 0 0;
		}
	}

	@keyframes dataFlow {
		0% {
			transform: translateX(-100%);
		}

		100% {
			transform: translateX(100%);
		}
	}

	#app>* {
		position: relative;
		z-index: 2;
	}

	button,
	a {
		transition: all 0.3s ease;
	}

	button:hover,
	a:hover {
		text-shadow: 0 0 8px var(--tech-blue);
		box-shadow: 0 0 15px rgba(0, 200, 255, 0.4);
	}

	/* 卡片效果 - 使用变量 */
	.card {
		background: var(--card-bg);
		border: 1px solid var(--card-border);
		border-radius: 8px;
		backdrop-filter: blur(10px);
		/* 增加磨砂感 */
		box-shadow:
			0 4px 20px var(--shadow-color),
			inset 0 1px 0 rgba(255, 255, 255, 0.1);
		/* 增加卡片文字颜色的适配 */
		color: var(--text-color);
	}

	/* --------------------------
     新增：切换按钮样式
     -------------------------- */
	.theme-switch-wrapper {
		position: fixed;
		top: 20px;
		right: 20px;
		z-index: 100;
		cursor: pointer;
	}

	.theme-switch {
		width: 60px;
		height: 30px;
		background: rgba(0, 0, 0, 0.5);
		border: 1px solid var(--tech-blue);
		border-radius: 30px;
		position: relative;
		transition: all 0.3s ease;
		box-shadow: 0 0 10px rgba(0, 170, 255, 0.2);
	}

	:root[data-theme="light"] .theme-switch {
		background: rgba(255, 255, 255, 0.8);
		border-color: var(--tech-blue);
	}

	.switch-handle {
		width: 24px;
		height: 24px;
		background: var(--tech-blue);
		border-radius: 50%;
		position: absolute;
		top: 2px;
		left: 3px;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 0.3s cubic-bezier(0.4, 0.0, 0.2, 1);
		box-shadow: 0 0 8px var(--tech-blue);
	}

	.theme-switch.active .switch-handle {
		left: 31px;
		background: #ffcc00;
		/* 太阳的颜色 */
		box-shadow: 0 0 8px #ffcc00;
	}

	.icon {
		width: 14px;
		height: 14px;
		color: #fff;
	}

	.sun {
		color: #fff;
		/* 太阳图标内部线条色 */
	}
</style>