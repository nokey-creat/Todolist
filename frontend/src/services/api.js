import api from '../utils/api';

// 用户认证服务
export const authService = {
  // 用户注册
  register: (userData) => {
    return api.post('/auth/register', userData);
  },
  
  // 用户登录
  login: (credentials) => {
    return api.post('/auth/login', credentials);
  }
};

// 任务服务
export const taskService = {
  // 获取所有任务
  getTasks: () => {
    return api.get('/tasks');
  },
  
  // 获取单个任务详情
  getTaskById: (taskId) => {
    return api.get(`/tasks/${taskId}`);
  },
  
  // 创建新任务
  createTask: (taskData) => {
    return api.post('/tasks', taskData);
  },
  
  // 更新任务
  updateTask: (taskId, taskData) => {
    return api.patch(`/tasks/${taskId}`, taskData);
  },
  
  // 删除任务
  deleteTask: (taskId) => {
    return api.delete(`/tasks/${taskId}`);
  },
  
  // 更改任务完成状态
  toggleTaskStatus: (taskId) => {
    return api.patch(`/tasks/${taskId}/completed`);
  }
};
