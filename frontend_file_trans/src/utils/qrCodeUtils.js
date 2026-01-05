import QRCode from 'qrcode';

// 生成二维码URL
export async function generateQrCodeUrl(text) {
  try {
    return await QRCode.toDataURL(text);
  } catch (error) {
    console.error('生成二维码失败:', error);
    return '';
  }
}