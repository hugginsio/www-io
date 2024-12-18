<script setup>
const props = defineProps({
    title: {
        type: Text,
        required: true,
    },
    tagline: {
        type: Text,
        required: false,
    },
});
const { title, tagline } = props;
</script>

<template>
    <div class="w-full py-20 rounded-md flex flex-col space-y-4">
        <span class="-ml-1 text-6xl font-bold tracking-tight">{{ props.title }}</span>
        <span>{{ props.tagline }}</span>
    </div>
</template>
