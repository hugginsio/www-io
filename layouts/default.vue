<template>
    <div class="h-full w-full colors-primary dark:colors-primary-invert" id="layout-default">
        <div class="max-w-3xl p-4 mx-auto">
            <div class="flex flex-col space-y-4">
                <AppHeader />
                <slot />
            </div>
        </div>
    </div>
</template>
