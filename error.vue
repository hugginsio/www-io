<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps({
    error: Object as () => NuxtError,
});
</script>

<template>
    <NuxtLayout>
        <Hero :title="`Error ${props.error?.statusCode}`" :tagline="props.error?.statusMessage" />
        <div>
            <div>
                <NuxtLink to="/">Go Home</NuxtLink>
            </div>
        </div>
    </NuxtLayout>
</template>
