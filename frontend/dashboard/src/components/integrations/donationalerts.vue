<script setup lang='ts'>
import { useDonationAlertsIntegration } from '@/api/index.js';
import IconDonationAlerts from '@/assets/icons/integrations/donationalerts.svg?component';
import OauthComponent from '@/components/integrations/variants/oauth.vue';

const manager = useDonationAlertsIntegration();
const { data } = manager.useData();
const logout = manager.useLogout();
const { data: authLink } = manager.useAuthLink();
</script>

<template>
	<oauth-component
		name="DonationAlerts"
		:data="data"
		:logout="() => logout.mutateAsync({})"
		:authLink="authLink?.link"
		:icon="IconDonationAlerts"
	/>
</template>
