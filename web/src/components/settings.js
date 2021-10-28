import { createClient, createMicrophoneAndCameraTracks } from "agora-rtc-react";

const appId = "agora-app-id";

export const config = { mode: "rtc", codec: "vp8", appId: appId };
export const useClient = createClient(config);
export const useMicrophoneAndCameraTracks = createMicrophoneAndCameraTracks();
