import axios from 'axios';
import type { AboutResponse } from './types';

const fallback: AboutResponse = {
  heroTitle: '启展跨境咨询服务',
  heroSubtitle: '为企业与家庭提供专业、合规、可落地的北美服务方案',
  companyOverview:
    '我们是一家专注于跨境咨询服务的团队，服务覆盖商业落地、移民规划、留学申请、税务协同与本地生活支持，强调长期陪跑与结果导向。',
  businesses: ['企业出海咨询', '加拿大移民服务', '留学与签证方案', '税务与财务协同', '本地安家支持'],
  advantages: ['一对一顾问机制', '透明流程与费用', '本地持牌合作网络', '多语种服务团队', '项目进度实时同步'],
  milestones: [
    { year: '2018', content: '团队成立，聚焦加拿大市场。' },
    { year: '2020', content: '服务升级，建立商业与家庭双线服务体系。' },
    { year: '2022', content: '上线数字化客户管理流程。' },
    { year: '2024', content: '拓展跨境企业综合服务能力。' },
  ],
  contact: {
    phone: '+1-604-123-4567',
    email: 'service@qizhan.ca',
    address: 'Vancouver, BC, Canada',
  },
};

export async function fetchAbout(): Promise<AboutResponse> {
  try {
    const { data } = await axios.get<AboutResponse>('/api/about');
    return data;
  } catch {
    return fallback;
  }
}
