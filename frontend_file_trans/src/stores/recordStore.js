// src/stores/recordStore.js
import { defineStore } from 'pinia'
import { ref, onMounted } from 'vue'

export const useRecordStore = defineStore('record', () => {
  const sendRecords = ref([])
  const pickupRecords = ref([])

  // 添加发送记录
  const addSendRecord = (fileName) => {
    const now = new Date()
    const record = {
      fileName: fileName,
      time: now.toLocaleTimeString('zh-CN'),
      timestamp: now.getTime(),
      type: 'send'
    }
    sendRecords.value.unshift(record)
    localStorage.setItem('sendRecords', JSON.stringify(sendRecords.value))
  }

  // 添加取件记录
  const addPickupRecord = (fileName) => {
    const now = new Date()
    const record = {
      fileName: fileName,
      time: now.toLocaleTimeString('zh-CN'),
      timestamp: now.getTime(),
      type: 'pickup'
    }
    pickupRecords.value.unshift(record)
    localStorage.setItem('pickupRecords', JSON.stringify(pickupRecords.value))
  }

  // 获取今日记录
  const getTodayRecords = (type) => {
    const today = new Date().toDateString()
    const records = type === 'send' ? sendRecords.value : pickupRecords.value
    return records.filter(record => {
      const recordDate = new Date(record.timestamp).toDateString()
      return recordDate === today
    })
  }

  // 初始化
  onMounted(() => {
    const savedSendRecords = localStorage.getItem('sendRecords')
    const savedPickupRecords = localStorage.getItem('pickupRecords')
    
    if (savedSendRecords) sendRecords.value = JSON.parse(savedSendRecords)
    if (savedPickupRecords) pickupRecords.value = JSON.parse(savedPickupRecords)
  })

  return {
    sendRecords,
    pickupRecords,
    addSendRecord,
    addPickupRecord,
    getTodayRecords
  }
})