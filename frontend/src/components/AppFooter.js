import React from 'react';
import { Layout, Typography } from 'antd';

const { Footer } = Layout;
const { Text } = Typography;

const AppFooter = () => {
  return (
    <Footer style={{ textAlign: 'center', background: '#f0f2f5', padding: '12px' }}>
      <Text type="secondary">Todo List App ©{new Date().getFullYear()} Created with React & Ant Design</Text>
    </Footer>
  );
};

export default AppFooter;
