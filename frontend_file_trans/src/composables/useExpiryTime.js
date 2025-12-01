// src/composables/useExpiryTime.js
import {
	ref,
	computed
} from 'vue'

export function useExpiryTime(initialDays = 1, initialUnit = '天') {
	const expiryDays = ref(initialDays)
	const timeUnit = ref(initialUnit)

	const minDays = ref(1)
	const maxDays = ref(30)

	// 计算总过期时间（毫秒）
	const totalExpiryTime = computed(() => {
		const days = expiryDays.value

		switch (timeUnit.value) {
			case '次':
				return days // 如果是次数，直接返回次数
			case '分钟':
				return days * 60 * 1000
			case '小时':
				return days * 60 * 60 * 1000
			case '天':
				return days * 24 * 60 * 60 * 1000
			default:
				return days * 24 * 60 * 60 * 1000
		}
	})

	const increaseDays = () => {
		if (expiryDays.value < maxDays.value) {
			expiryDays.value++
		}
	}

	const decreaseDays = () => {
		if (expiryDays.value > minDays.value) {
			expiryDays.value--
		}
	}

	const setExpiryDays = (days) => {
		if (days >= minDays.value && days <= maxDays.value) {
			expiryDays.value = days
		}
	}

	const setTimeUnit = (unit) => {
		const validUnits = ['次', '分钟', '小时', '天']
		if (validUnits.includes(unit)) {
			timeUnit.value = unit
		}
	}

	return {
		expiryDays,
		timeUnit,
		minDays,
		maxDays,
		totalExpiryTime,
		increaseDays,
		decreaseDays,
		setExpiryDays,
		setTimeUnit
	}
}