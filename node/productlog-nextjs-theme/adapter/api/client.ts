import axios, { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";

class apiClient {
    timeout: number;

    constructor(timeout: number) {
        this.timeout = timeout;
    }

    call(endpoint: Endpoint, req: CustomRequest) {
        const options: AxiosRequestConfig = {
            url: apiRoot + endpoint.url,
            method: endpoint.method,
        };
        axios(options)
            .then((res: AxiosResponse<USER[]>) => {
                const { data, status } = res;
                //やりたいことをやる
            })
            .catch((e: AxiosError<{ error: string }>) => {
                // エラー処理
                console.log(e.message);
            });

    }
}