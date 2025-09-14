import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { Card, Button, Typography, Descriptions, Spin, Popconfirm, Modal, Form, Input, DatePicker, Tag, message } from 'antd';
import { EditOutlined, DeleteOutlined, ArrowLeftOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons';
import { taskService } from '../services/api';
import { formatDateForAPI, formatDateForDisplay, isDeadlinePassed } from '../utils/dateUtils';
import dayjs from 'dayjs';

const { Title, Text } = Typography;
const { TextArea } = Input;

const TaskDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [task, setTask] = useState(null);
  const [loading, setLoading] = useState(true);
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [submitting, setSubmitting] = useState(false);

  // 获取任务详情
  const fetchTaskDetail = async () => {
    setLoading(true);
    try {
      const response = await taskService.getTaskById(id);
      setTask(response.data);
      
      // 设置表单初始值
      form.setFieldsValue({
        title: response.data.Title,
        description: response.data.Description,
        deadline: dayjs(response.data.Deadline),
      });
    } catch (error) {
      if (error.response && error.response.status === 401) {
        message.error('没有权限查看该任务或任务不存在');
        navigate('/tasks');
      } else {
        console.error('获取任务详情失败:', error);
        message.error('获取任务详情失败');
      }
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTaskDetail();
  }, [id]);

  // 切换任务完成状态
  const handleToggleStatus = async () => {
    try {
      await taskService.toggleTaskStatus(id);
      message.success('任务状态已更新');
      fetchTaskDetail();
    } catch (error) {
      console.error('更新任务状态失败:', error);
      message.error('更新任务状态失败');
    }
  };

  // 删除任务
  const handleDeleteTask = async () => {
    try {
      await taskService.deleteTask(id);
      message.success('任务已删除');
      navigate('/tasks');
    } catch (error) {
      console.error('删除任务失败:', error);
      message.error('删除任务失败');
    }
  };

  // 更新任务
  const handleUpdateTask = async (values) => {
    setSubmitting(true);
    try {
      const taskData = {
        ...values,
        deadline: formatDateForAPI(values.deadline),
      };
      
      await taskService.updateTask(id, taskData);
      message.success('任务更新成功！');
      setEditModalVisible(false);
      fetchTaskDetail();
    } catch (error) {
      console.error('更新任务失败:', error);
      message.error('更新任务失败，请重试');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '400px' }}>
        <Spin size="large" />
      </div>
    );
  }

  if (!task) {
    return null;
  }

  // 计算创建时间显示
  const createdAtDisplay = task.CreatedAt ? dayjs(task.CreatedAt).format('YYYY-MM-DD HH:mm:ss') : '未知';
  const updatedAtDisplay = task.UpdatedAt ? dayjs(task.UpdatedAt).format('YYYY-MM-DD HH:mm:ss') : '未知';

  return (
    <div style={{ maxWidth: 800, margin: '0 auto', padding: '20px' }}>
      <Button 
        icon={<ArrowLeftOutlined />} 
        onClick={() => navigate('/tasks')}
        style={{ marginBottom: 16 }}
      >
        返回任务列表
      </Button>

      <Card
        title={
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <Title level={3}>{task.Title}</Title>
            <Tag 
              color={task.Completed ? 'success' : isDeadlinePassed(task.Deadline) ? 'error' : 'processing'}
              style={{ marginRight: 0, fontSize: '14px', padding: '0 10px' }}
            >
              {task.Completed ? '已完成' : isDeadlinePassed(task.Deadline) ? '已过期' : '进行中'}
            </Tag>
          </div>
        }
        extra={
          <div>
            <Button 
              icon={task.Completed ? <CloseCircleOutlined /> : <CheckCircleOutlined />}
              onClick={handleToggleStatus}
              style={{ marginRight: 8 }}
            >
              {task.Completed ? '标记为未完成' : '标记为已完成'}
            </Button>
            <Button 
              type="primary" 
              icon={<EditOutlined />} 
              onClick={() => setEditModalVisible(true)}
              style={{ marginRight: 8 }}
            >
              编辑
            </Button>
            <Popconfirm
              title="确定要删除这个任务吗？"
              onConfirm={handleDeleteTask}
              okText="确定"
              cancelText="取消"
              placement="topRight"
            >
              <Button danger icon={<DeleteOutlined />}>删除</Button>
            </Popconfirm>
          </div>
        }
      >
        <Descriptions column={1} bordered>
          <Descriptions.Item label="截止日期">
            {formatDateForDisplay(task.Deadline)}
          </Descriptions.Item>
          <Descriptions.Item label="任务描述">
            <Text style={{ whiteSpace: 'pre-wrap' }}>{task.Description}</Text>
          </Descriptions.Item>
          <Descriptions.Item label="创建时间">
            {createdAtDisplay}
          </Descriptions.Item>
          <Descriptions.Item label="最后更新">
            {updatedAtDisplay}
          </Descriptions.Item>
        </Descriptions>
      </Card>

      {/* 编辑任务的表单模态框 */}
      <Modal
        title="编辑任务"
        open={editModalVisible}
        onCancel={() => setEditModalVisible(false)}
        footer={null}
      >
        <Form
          form={form}
          layout="vertical"
          onFinish={handleUpdateTask}
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
            />
          </Form.Item>

          <Form.Item>
            <div style={{ display: 'flex', justifyContent: 'flex-end', gap: 8 }}>
              <Button onClick={() => setEditModalVisible(false)}>
                取消
              </Button>
              <Button type="primary" htmlType="submit" loading={submitting}>
                更新
              </Button>
            </div>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default TaskDetail;
