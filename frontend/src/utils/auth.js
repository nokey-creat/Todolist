// 检查是否已认证
export const isAuthenticated = () => {
  const token = localStorage.getItem('token');
  return !!token;
};

// 获取JWT令牌
export const getToken = () => {
  return localStorage.getItem('token');
};

// 设置JWT令牌
export const setToken = (token) => {
  localStorage.setItem('token', token);
};

// 清除JWT令牌
export const clearToken = () => {
  localStorage.removeItem('token');
};
