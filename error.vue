<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps({
    error: Object as () => NuxtError,
});
</script>

<template>
    <NuxtLayout>
        <Hero>
            <span class="-ml-1 text-6xl font-bold tracking-tight">Error {{ props.error?.statusCode }}</span>
            <span>{{ props.error?.statusMessage }}</span>
        </Hero>
        <div>
            <div>
                <NuxtLink to="/">Go Home</NuxtLink>
            </div>
        </div>
    </NuxtLayout>
</template>
