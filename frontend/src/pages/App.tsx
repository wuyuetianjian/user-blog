import { useEffect, useState } from 'react';
import {
  CheckCircleOutlined,
  EnvironmentOutlined,
  MailOutlined,
  PhoneOutlined,
} from '@ant-design/icons';
import {
  FooterToolbar,
  PageContainer,
  ProCard,
  ProLayout,
  StatisticCard,
} from '@ant-design/pro-components';
import { Alert, Col, Divider, Row, Space, Spin, Tag, Timeline, Typography } from 'antd';
import { fetchAbout } from '../api';
import type { AboutResponse } from '../types';

const { Title, Paragraph, Text } = Typography;

export default function App() {
  const [data, setData] = useState<AboutResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchAbout()
      .then(setData)
      .finally(() => setLoading(false));
  }, []);

  if (loading || !data) {
    return (
      <div className="center">
        <Spin size="large" />
      </div>
    );
  }

  return (
    <ProLayout
      title="启展咨询"
      layout="top"
      fixedHeader
      contentWidth="Fixed"
      headerTitleRender={() => <span>启展咨询 · 关于我们</span>}
      menuDataRender={() => [{ path: '/', name: '关于我们' }]}
    >
      <PageContainer ghost>
        <ProCard className="hero-card" bordered>
          <Title>{data.heroTitle}</Title>
          <Paragraph>{data.heroSubtitle}</Paragraph>
          <Space wrap>
            {data.businesses.map((item) => (
              <Tag key={item} color="blue">
                {item}
              </Tag>
            ))}
          </Space>
        </ProCard>

        <Row gutter={[16, 16]}>
          <Col xs={24} lg={16}>
            <ProCard title="公司介绍" bordered>
              <Paragraph>{data.companyOverview}</Paragraph>
            </ProCard>
          </Col>
          <Col xs={24} lg={8}>
            <StatisticCard.Group direction="column">
              <StatisticCard statistic={{ title: '累计服务客户', value: 1200, suffix: '+' }} />
              <StatisticCard statistic={{ title: '顾问与合作专家', value: 40, suffix: '+' }} />
            </StatisticCard.Group>
          </Col>
        </Row>

        <Divider />

        <ProCard title="我们的优势" bordered>
          <Row gutter={[12, 12]}>
            {data.advantages.map((item) => (
              <Col xs={24} md={12} key={item}>
                <Alert message={item} type="success" icon={<CheckCircleOutlined />} showIcon />
              </Col>
            ))}
          </Row>
        </ProCard>

        <Divider />

        <ProCard title="发展里程碑" bordered>
          <Timeline
            items={data.milestones.map((m) => ({
              children: (
                <>
                  <Text strong>{m.year}</Text>
                  <br />
                  <Text>{m.content}</Text>
                </>
              ),
            }))}
          />
        </ProCard>
      </PageContainer>

      <FooterToolbar>
        <Space direction="vertical" size={2}>
          <Space>
            <PhoneOutlined />
            <Text>{data.contact.phone}</Text>
          </Space>
          <Space>
            <MailOutlined />
            <Text>{data.contact.email}</Text>
          </Space>
          <Space>
            <EnvironmentOutlined />
            <Text>{data.contact.address}</Text>
          </Space>
        </Space>
      </FooterToolbar>
    </ProLayout>
  );
}
