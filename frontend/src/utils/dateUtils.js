import dayjs from 'dayjs';

// 格式化日期为后端需要的格式
export const formatDateForAPI = (date) => {
  return dayjs(date).format('YYYY-MM-DDT00:00:00Z');
};

// 格式化日期为显示格式
export const formatDateForDisplay = (dateString) => {
  return dayjs(dateString).format('YYYY-MM-DD');
};

// 检查是否过期
export const isDeadlinePassed = (deadline) => {
  return dayjs().isAfter(dayjs(deadline));
};
