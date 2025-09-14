import React, { useState, useEffect } from 'react';
import { List, Card, Tag, Button, Modal, Form, Input, DatePicker, Typography, Spin, Empty, message } from 'antd';
import { PlusOutlined, CheckCircleOutlined, CloseCircleOutlined, CalendarOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';
import { taskService } from '../services/api';
import { formatDateForAPI, formatDateForDisplay, isDeadlinePassed } from '../utils/dateUtils';
import dayjs from 'dayjs';

const { Title, Text } = Typography;
const { TextArea } = Input;

const TaskList = () => {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [modalVisible, setModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);
  const navigate = useNavigate();

  // 排序任务：未完成的排在前面，在各自分组内按ID降序排列
  const sortTasks = (taskList) => {
    return taskList.sort((a, b) => {
      // 首先按完成状态排序：未完成(false)在前，已完成(true)在后
      if (a.Completed !== b.Completed) {
        return a.Completed - b.Completed;
      }
      // 在同一完成状态内，按ID降序排列（ID大的在前）
      return b.ID - a.ID;
    });
  };

  // 获取所有任务
  const fetchTasks = async () => {
    setLoading(true);
    try {
      const response = await taskService.getTasks();
      const sortedTasks = sortTasks(response.data);
      setTasks(sortedTasks);
    } catch (error) {
      console.error('获取任务失败:', error);
      message.error('获取任务列表失败');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  // 创建新任务
  const handleCreateTask = async (values) => {
    setSubmitting(true);
    try {
      const taskData = {
        ...values,
        deadline: formatDateForAPI(values.deadline),
      };
      
      await taskService.createTask(taskData);
      message.success('任务创建成功！');
      setModalVisible(false);
      form.resetFields();
      fetchTasks();
    } catch (error) {
      console.error('创建任务失败:', error);
      message.error('创建任务失败，请重试');
    } finally {
      setSubmitting(false);
    }
  };

  // 切换任务完成状态
  const handleToggleStatus = async (taskId, e) => {
    e.stopPropagation();
    try {
      await taskService.toggleTaskStatus(taskId);
      message.success('任务状态已更新');
      // 重新获取并排序任务
      fetchTasks();
    } catch (error) {
      console.error('更新任务状态失败:', error);
      message.error('更新任务状态失败');
    }
  };

  // 处理点击任务，导航到详情页
  const handleTaskClick = (taskId) => {
    navigate(`/tasks/${taskId}`);
  };

  return (
    <div style={{ maxWidth: 1000, margin: '0 auto', padding: '20px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 24 }}>
        <Title level={2}>我的待办事项</Title>
        <Button 
          type="primary" 
          icon={<PlusOutlined />} 
          onClick={() => setModalVisible(true)}
        >
          新建任务
        </Button>
      </div>

      {loading ? (
        <div style={{ textAlign: 'center', margin: '40px 0' }}>
          <Spin size="large" />
        </div>
      ) : tasks.length === 0 ? (
        <Empty description="暂无待办事项" />
      ) : (
        <List
          grid={{ 
            gutter: 16, 
            xs: 1, 
            sm: 1, 
            md: 2, 
            lg: 3, 
            xl: 3, 
            xxl: 4 
          }}
          dataSource={tasks}
          renderItem={(task) => (
            <List.Item>
              <Card 
                hoverable
                onClick={() => handleTaskClick(task.ID)}
                title={
                  <div style={{ 
                    display: 'flex', 
                    alignItems: 'center', 
                    justifyContent: 'space-between',
                    maxWidth: '100%',
                    overflow: 'hidden'
                  }}>
                    <Text 
                      ellipsis 
                      style={{ 
                        maxWidth: '80%',
                        textDecoration: task.Completed ? 'line-through' : 'none',
                        color: task.Completed ? '#999' : 'inherit',
                      }}
                    >
                      {task.Title}
                    </Text>
                    <Tag 
                      color={task.Completed ? 'success' : isDeadlinePassed(task.Deadline) ? 'error' : 'processing'}
                      onClick={(e) => handleToggleStatus(task.ID, e)}
                      style={{ cursor: 'pointer' }}
                    >
                      {task.Completed ? '已完成' : isDeadlinePassed(task.Deadline) ? '已过期' : '进行中'}
                    </Tag>
                  </div>
                }
                actions={[
                  <div key="deadline">
                    <CalendarOutlined /> {formatDateForDisplay(task.Deadline)}
                  </div>,
                  <Button 
                    key="status" 
                    type="text"
                    icon={task.Completed ? <CloseCircleOutlined /> : <CheckCircleOutlined />}
                    onClick={(e) => handleToggleStatus(task.ID, e)}
                  >
                    {task.Completed ? '标记未完成' : '标记完成'}
                  </Button>
                ]}
              >
                <Card.Meta
                  description={
                    <Text ellipsis={{ rows: 2 }} style={{ color: '#666' }}>
                      {task.Description}
                    </Text>
                  }
                />
              </Card>
            </List.Item>
          )}
        />
      )}

      {/* 创建任务的表单模态框 */}
      <Modal
        title="创建新任务"
        open={modalVisible}
        onCancel={() => setModalVisible(false)}
        footer={null}
      >
        <Form
          form={form}
          layout="vertical"
          onFinish={handleCreateTask}
        >
          <Form.Item
            name="title"
            label="任务标题"
            rules={[{ required: true, message: '请输入任务标题' }]}
          >
            <Input placeholder="请输入任务标题" maxLength={100} />
          </Form.Item>

          <Form.Item
            name="description"
            label="任务描述"
            rules={[{ required: true, message: '请输入任务描述' }]}
          >
            <TextArea 
              placeholder="请输入任务描述" 
              rows={4} 
              maxLength={1000}
              showCount 
            />
          </Form.Item>

          <Form.Item
            name="deadline"
            label="截止日期"
            rules={[{ required: true, message: '请选择截止日期' }]}
          >
            <DatePicker 
              style={{ width: '100%' }} 
              placeholder="选择截止日期"
              format="YYYY-MM-DD"
              disabledDate={(current) => current && current < dayjs().startOf('day')}
            />
          </Form.Item>

          <Form.Item>
            <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 8 }}>
              <Button onClick={() => setModalVisible(false)}>
                取消
              </Button>
              <Button type="primary" htmlType="submit" loading={submitting}>
                创建
              </Button>
            </div>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default TaskList;
