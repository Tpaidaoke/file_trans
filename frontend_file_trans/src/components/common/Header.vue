<!-- src/components/common/Header.vue -->
<template>
	<div class="header">
		<img class="logo" :src="currentLogo" alt="logo" @click="switchLogo">
		<div class="title">{{ title }}</div>
	</div>
</template>

<script setup>
	import {
		ref
	} from 'vue'

	const props = defineProps({
		title: {
			type: String,
			default: '发取文件'
		}
	})

	// 图片轮换
	const logos = [
		"src/assets/send_logo.png",
		"src/assets/view_logo.png",
		"src/assets/pickup_logo.png"
	]

	const currentLogoIndex = ref(0)
	const currentLogo = ref(logos[0])

	// 点击更换照片
	const switchLogo = () => {
		currentLogoIndex.value = (currentLogoIndex.value + 1) % logos.length
		currentLogo.value = logos[currentLogoIndex.value]
	}
</script>

<style scoped>
	.header {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-bottom: 30px;
	}

	.logo {
		width: 32px;
		height: 32px;
		margin-right: 12px;
		cursor: pointer;
	}

	.title {
		font-size: 20px;
		font-weight: 600;
		color: #00007f;
	}
</style>