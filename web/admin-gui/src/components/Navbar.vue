<script setup lang="ts">
import { useAuthStore } from '@/stores/AuthStore';
import { deleteCookie } from '@/utils/cookies';
import { RouterLink, useRouter } from 'vue-router'

const router = useRouter()

function logout() {
	deleteCookie('JWTtoken')
	deleteCookie('JWTexpire')
	useAuthStore().clearAll()
	router.push('/')
}
</script>

<template>
	<header>
		<h2><span style="color:var(--brand-blue)">Go</span>hCMS</h2>
		<nav>
			<RouterLink to="/home">Accueil</RouterLink>
			<RouterLink to="/articles">Articles</RouterLink>
			<a class="disconnect-link" @click="logout()">Se déconnecter</a>
		</nav>
	</header>
</template>

<style scoped>
header {
	display: flex;
	align-items: center;
	padding: 0 24px;

	box-shadow: #0005 0 0 10px;
	
	width: 100%;
}

header h2 {
	margin: 0;
	padding: 10px 0;

	width: fit-content;

	color: var(--neutral-dark);
}

header h2 span {
	font-family: inherit;
}

header nav {
	display: flex;
	gap: 16px;

	width: 100%;
	padding: 20px;
}

header nav a {
	text-decoration: none;
	border-radius: var(--radius);
	padding: 4px 16px;
	cursor: pointer;
	transition: 150ms background-color;
}

header nav a:hover {
	background-color: var(--neutral-verylight);
}

.disconnect-link {
	margin-left: auto;
}

.disconnect-link:hover {
	background-color: #fff;
	color: var(--error);
}
</style>