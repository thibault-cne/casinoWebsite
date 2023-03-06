import { getCurrentInstance } from 'vue'

export default function useEmitter() {
    const internalInstance = getCurrentInstance(); 

    if (internalInstance != null) {
        const emitter = internalInstance.appContext.config.globalProperties.emitter;

        return emitter;
    }

    return null;
}