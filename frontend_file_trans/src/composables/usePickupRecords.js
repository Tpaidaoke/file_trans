// src/composables/usePickupRecords.js
import { ref, onMounted } from 'vue'

export function usePickupRecords() {
  const pickupRecords = ref([])

  const addPickupRecord = (fileName) => {
    const now = new Date()
    const record = {
      fileName: fileName,
      time: now.toLocaleTimeString('zh-CN', { 
        hour: '2-digit', 
        minute: '2-digit', 
        second: '2-digit' 
      }),
      timestamp: now.getTime()
    }
    pickupRecords.value.unshift(record)
    
    // 保存到本地存储
    localStorage.setItem('pickupRecords', JSON.stringify(pickupRecords.value))
  }

  // 初始化：从本地存储加载取件记录
  onMounted(() => {
    const savedRecords = localStorage.getItem('pickupRecords')
    if (savedRecords) {
      pickupRecords.value = JSON.parse(savedRecords)
    }
  })

  return {
    pickupRecords,
    addPickupRecord
  }
}