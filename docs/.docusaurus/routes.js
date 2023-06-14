import React from 'react';
import ComponentCreator from '@docusaurus/ComponentCreator';

export default [
  {
    path: '/interchain-security/',
    component: ComponentCreator('/interchain-security/', '4a2'),
    routes: [
      {
        path: '/interchain-security/',
        component: ComponentCreator('/interchain-security/', '725'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/adrs/adr-001-key-assignment',
        component: ComponentCreator('/interchain-security/adrs/adr-001-key-assignment', 'b01'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/adrs/adr-002-throttle',
        component: ComponentCreator('/interchain-security/adrs/adr-002-throttle', '049'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/adrs/adr-003-equivocation-gov-proposal',
        component: ComponentCreator('/interchain-security/adrs/adr-003-equivocation-gov-proposal', 'd6f'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/adrs/adr-template',
        component: ComponentCreator('/interchain-security/adrs/adr-template', 'b11'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/adrs/intro',
        component: ComponentCreator('/interchain-security/adrs/intro', '130'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/consumer-development/app-integration',
        component: ComponentCreator('/interchain-security/consumer-development/app-integration', '634'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/consumer-development/consumer-chain-governance',
        component: ComponentCreator('/interchain-security/consumer-development/consumer-chain-governance', 'f51'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/consumer-development/consumer-chain-upgrade-procedure',
        component: ComponentCreator('/interchain-security/consumer-development/consumer-chain-upgrade-procedure', '776'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/consumer-development/offboarding',
        component: ComponentCreator('/interchain-security/consumer-development/offboarding', 'b9a'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/consumer-development/onboarding',
        component: ComponentCreator('/interchain-security/consumer-development/onboarding', '81f'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/faq',
        component: ComponentCreator('/interchain-security/faq', 'e46'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/features/key-assignment',
        component: ComponentCreator('/interchain-security/features/key-assignment', '6a8'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/features/proposals',
        component: ComponentCreator('/interchain-security/features/proposals', '85e'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/features/reward-distribution',
        component: ComponentCreator('/interchain-security/features/reward-distribution', '133'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/features/slashing',
        component: ComponentCreator('/interchain-security/features/slashing', 'e4f'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/introduction/overview',
        component: ComponentCreator('/interchain-security/introduction/overview', '287'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/introduction/params',
        component: ComponentCreator('/interchain-security/introduction/params', 'aae'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/introduction/technical-specification',
        component: ComponentCreator('/interchain-security/introduction/technical-specification', 'e7d'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/introduction/terminology',
        component: ComponentCreator('/interchain-security/introduction/terminology', '11e'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/validators/joining-testnet',
        component: ComponentCreator('/interchain-security/validators/joining-testnet', '82d'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/validators/overview',
        component: ComponentCreator('/interchain-security/validators/overview', 'ccb'),
        exact: true,
        sidebar: "tutorialSidebar"
      },
      {
        path: '/interchain-security/validators/withdraw_rewards',
        component: ComponentCreator('/interchain-security/validators/withdraw_rewards', '2b7'),
        exact: true,
        sidebar: "tutorialSidebar"
      }
    ]
  },
  {
    path: '*',
    component: ComponentCreator('*'),
  },
];
