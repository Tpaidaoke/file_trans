// src/composables/useSendRecords.js
import { ref, onMounted, computed } from 'vue'

export function useSendRecords() {
  const sendRecords = ref([])

  const addSendRecord = (fileName) => {
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
    sendRecords.value.unshift(record)
    
    // 保存到本地存储
    localStorage.setItem('sendRecords', JSON.stringify(sendRecords.value))
  }

  // 获取今日发送记录
  const todaySendRecords = computed(() => {
    const today = new Date().toDateString()
    return sendRecords.value.filter(record => {
      const recordDate = new Date(record.timestamp).toDateString()
      return recordDate === today
    })
  })

  // 初始化：从本地存储加载发件记录
  onMounted(() => {
    const savedRecords = localStorage.getItem('sendRecords')
    if (savedRecords) {
      sendRecords.value = JSON.parse(savedRecords)
    }
  })

  return {
    sendRecords,
    todaySendRecords,
    addSendRecord
  }
}