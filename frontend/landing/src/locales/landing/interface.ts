import type { BotFeature } from '@/data/landing/botFeatures.js';
import type { PricingPlansLocale } from '@/data/landing/pricingPlans.js';
import type { NavMenuLocale } from '@/data/landing/sections.js';
import type { TeamMemberLocale } from '@/data/landing/team.js';

interface ILandingLocale {
  navMenu: NavMenuLocale[];
  tagline: string;
  buttons: {
    startForFree: string;
    getStarted: string;
    buyPlan: string;
    learnMore: string;
    login: string;
    tryFeature: string;
  };
  sections: {
    firstScreen: {
      title: string;
    };
    statLine: {
      statPlaceholder: string;
    };
    features: {
      title: string;
      featuresInDev: string;
      content: BotFeature[];
    };
    integrations: {
      preTitle: string;
      title: string;
      description: string;
    };
    reviews: {
      title: string;
    };
    team: {
      title: string;
      description: string;
      founder: string;
      members: TeamMemberLocale;
    };
    pricing: {
      title: string;
      features: string;
      perMonth: string;
      plans: PricingPlansLocale;
    };
    subscribeForUpdates: {
      title: string;
      description: string;
      inputPlaceholder: string;
    };
    footer: {
      rights: string;
    };
  };
}

export default ILandingLocale;
