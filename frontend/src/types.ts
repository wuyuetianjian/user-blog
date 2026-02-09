export interface AboutResponse {
  heroTitle: string;
  heroSubtitle: string;
  companyOverview: string;
  businesses: string[];
  advantages: string[];
  milestones: Array<{
    year: string;
    content: string;
  }>;
  contact: {
    phone: string;
    email: string;
    address: string;
  };
}
