import React from 'react';
import { Layout, Menu, Button } from 'antd';
import { HomeOutlined, LogoutOutlined, UserOutlined, UnorderedListOutlined } from '@ant-design/icons';
import { Link, useLocation } from 'react-router-dom';

const { Header } = Layout;

const AppHeader = ({ loggedIn, onLogout }) => {
  const location = useLocation();
  
  return (
    <Header style={{ 
      background: '#fff', 
      boxShadow: '0 2px 8px rgba(0, 0, 0, 0.06)',
      display: 'flex', 
      justifyContent: 'space-between',
      alignItems: 'center',
      padding: '0 20px'
    }}>
      <div style={{ display: 'flex', alignItems: 'center' }}>
        <div style={{ 
          fontSize: '20px', 
          fontWeight: 'bold', 
          marginRight: '40px',
          color: '#1890ff'
        }}>
          <Link to="/" style={{ color: 'inherit' }}>
            <UnorderedListOutlined /> Todo List
          </Link>
        </div>
        
        {loggedIn && (
          <Menu
            mode="horizontal"
            selectedKeys={[location.pathname]}
            style={{ border: 'none' }}
          >
            <Menu.Item key="/tasks" icon={<HomeOutlined />}>
              <Link to="/tasks">我的任务</Link>
            </Menu.Item>
          </Menu>
        )}
      </div>
      
      <div>
        {loggedIn ? (
          <Button 
            icon={<LogoutOutlined />}
            onClick={onLogout}
          >
            退出登录
          </Button>
        ) : (
          <Menu 
            mode="horizontal" 
            selectedKeys={[location.pathname]}
            style={{ border: 'none' }}
          >
            <Menu.Item key="/login" icon={<UserOutlined />}>
              <Link to="/login">登录</Link>
            </Menu.Item>
            <Menu.Item key="/register" icon={<UserOutlined />}>
              <Link to="/register">注册</Link>
            </Menu.Item>
          </Menu>
        )}
      </div>
    </Header>
  );
};

export default AppHeader;
