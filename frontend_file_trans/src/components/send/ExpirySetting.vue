<!-- src/components/send/ExpirySetting.vue -->
<template>
	<div class="expiry-setting">
		<h3 class="subsection-title">过期设置</h3>
		<div class="setting-content">
			<div class="number-input">
				<button class="btn-minus" @click="decreaseDays" :disabled="expiryDays <= minDays">-</button>
				<span class="number">{{ expiryDays }}</span>
				<button class="btn-plus" @click="increaseDays" :disabled="expiryDays >= maxDays">+</button>
			</div>
			<select v-model="timeUnit" class="unit-select" @change="handleUnitChange">
				<option value="次">次</option>
				<option value="分钟">分钟</option>
				<option value="小时">小时</option>
				<option value="天">天</option>
			</select>
		</div>
	</div>
</template>

<script setup>
	import {
		ref,
		watch
	} from 'vue'

	const props = defineProps({
		expiryDays: {
			type: Number,
			default: 1
		},
		timeUnit: {
			type: String,
			default: '天'
		},
		minDays: {
			type: Number,
			default: 1
		},
		maxDays: {
			type: Number,
			default: 30
		}
	})

	const emit = defineEmits(['update:expiryDays', 'update:timeUnit'])

	const expiryDays = ref(props.expiryDays)
	const timeUnit = ref(props.timeUnit)

	// 监听内部变化并通知父组件
	watch(expiryDays, (newVal) => {
		emit('update:expiryDays', newVal)
	})

	watch(timeUnit, (newVal) => {
		emit('update:timeUnit', newVal)
	})

	// 监听父组件的变化
	watch(() => props.expiryDays, (newVal) => {
		expiryDays.value = newVal
	})

	watch(() => props.timeUnit, (newVal) => {
		timeUnit.value = newVal
	})

	const increaseDays = () => {
		if (expiryDays.value < props.maxDays) {
			expiryDays.value++
		}
	}

	const decreaseDays = () => {
		if (expiryDays.value > props.minDays) {
			expiryDays.value--
		}
	}

	const handleUnitChange = () => {
		// 当单位改变时，可以重置天数或进行其他逻辑
		console.log('时间单位改变为:', timeUnit.value)
	}
</script>

<style scoped>
	.expiry-setting {
		flex: 1;
	}

	.subsection-title {
		font-size: 16px;
		font-weight: 500;
		margin-bottom: 12px;
		color: #1a1a1a;
	}

	.setting-content {
		display: flex;
		align-items: center;
		gap: 12px;
	}

	.number-input {
		display: flex;
		align-items: center;
		border: 1px solid #d9d9d9;
		border-radius: 6px;
		overflow: hidden;
	}

	.number-input button {
		width: 32px;
		height: 32px;
		border: none;
		background: #f5f5f5;
		cursor: pointer;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: background 0.3s;
	}

	.number-input button:hover:not(:disabled) {
		background: #e6f7ff;
	}

	.number-input button:disabled {
		background: #f5f5f5;
		color: #ccc;
		cursor: not-allowed;
	}

	.number {
		width: 40px;
		text-align: center;
		font-weight: 500;
		font-size: 14px;
	}

	.unit-select {
		padding: 6px 12px;
		border: 1px solid #d9d9d9;
		border-radius: 6px;
		background-color: #fff;
		color: #666;
		font-size: 14px;
		cursor: pointer;
		transition: border-color 0.3s;
		appearance: none;
		background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 24 24' fill='none' stroke='%23666' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpolyline points='6 9 12 15 18 9'%3E%3C/polyline%3E%3C/svg%3E");
		background-repeat: no-repeat;
		background-position: right 8px center;
		padding-right: 30px;
	}

	.unit-select:hover {
		border-color: #1890ff;
	}

	.unit-select:focus {
		outline: none;
		border-color: #1890ff;
	}

	/* 响应式设计 */
	@media (max-width: 768px) {
		.setting-content {
			flex-direction: column;
			align-items: stretch;
			gap: 8px;
		}

		.number-input {
			justify-content: center;
		}

		.unit-select {
			text-align: center;
		}
	}
</style>