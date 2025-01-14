<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps({
    error: Object as () => NuxtError,
});
</script>

<template>
    <NuxtLayout name="card">
        <div class="prose">
            <p>Error {{ props.error?.statusCode }}</p>
            <div class="gray">
                <p>{{ props.error?.statusMessage }}</p>
                <NuxtLink to="/">Go Home</NuxtLink>
            </div>
        </div>
    </NuxtLayout>
</template>
