<script setup lang="ts">
import type { NuxtError } from "#app";

const props = defineProps({
    error: Object as () => NuxtError,
});
</script>

<template>
    <NuxtLayout>
        <div class="flex flex-col space-y-4">
            <div class="w-full flex flex-col space-y-4 font-bold">
                <span class="text-5xl font-bold">Error {{ props.error?.statusCode }}</span>
                <span class="p-3 rounded colors-primary-invert flex items-center uppercase font-mono space-x-6">
                    <NuxtLink to="/">Home</NuxtLink>
                </span>
            </div>
            <div>
                <p class="prose">{{ props.error?.statusMessage }}</p>
            </div>
        </div>
    </NuxtLayout>
</template>
