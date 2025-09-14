import React, { useState, useEffect } from 'react';
import { Route, Routes, Navigate, useNavigate } from 'react-router-dom';
import { Layout, message, ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';

import Login from './pages/Login';
import Register from './pages/Register';
import TaskList from './pages/TaskList';
import TaskDetail from './pages/TaskDetail';
import NotFound from './pages/NotFound';
import AppHeader from './components/AppHeader';
import AppFooter from './components/AppFooter';
import { isAuthenticated } from './utils/auth';

const { Content } = Layout;

// 受保护的路由
const ProtectedRoute = ({ children }) => {
  if (!isAuthenticated()) {
    message.error('您需要先登录');
    return <Navigate to="/login" replace />;
  }
  return children;
};

function App() {
  const [loggedIn, setLoggedIn] = useState(isAuthenticated());
  const navigate = useNavigate();

  useEffect(() => {
    // 监听登录状态
    const checkAuth = () => {
      setLoggedIn(isAuthenticated());
    };
    
    checkAuth();
    window.addEventListener('storage', checkAuth);
    
    return () => {
      window.removeEventListener('storage', checkAuth);
    };
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    setLoggedIn(false);
    navigate('/login');
    message.success('已成功退出登录');
  };

  return (
    <ConfigProvider locale={zhCN}>
      <Layout className="page-container">
        <AppHeader loggedIn={loggedIn} onLogout={handleLogout} />
        <Content className="site-layout-content">
          <Routes>
            <Route path="/" element={
              loggedIn ? <Navigate to="/tasks" replace /> : <Navigate to="/login" replace />
            } />
            <Route path="/login" element={
              loggedIn ? <Navigate to="/tasks" replace /> : <Login onLoginSuccess={() => setLoggedIn(true)} />
            } />
            <Route path="/register" element={
              loggedIn ? <Navigate to="/tasks" replace /> : <Register />
            } />
            <Route path="/tasks" element={
              <ProtectedRoute>
                <TaskList />
              </ProtectedRoute>
            } />
            <Route path="/tasks/:id" element={
              <ProtectedRoute>
                <TaskDetail />
              </ProtectedRoute>
            } />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Content>
        <AppFooter />
      </Layout>
    </ConfigProvider>
  );
}

export default App;
