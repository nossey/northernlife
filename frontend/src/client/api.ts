// tslint:disable
/**
 * Swagger Example API
 * This is a sample server celler server.
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as globalImportUrl from 'url';
import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface ModelErrorMessage
 */
export interface ModelErrorMessage {
    /**
     * 
     * @type {string}
     * @memberof ModelErrorMessage
     */
    message?: string;
}
/**
 * 
 * @export
 * @interface ModelFileImageUploadModel
 */
export interface ModelFileImageUploadModel {
    /**
     * 
     * @type {string}
     * @memberof ModelFileImageUploadModel
     */
    image: string;
}
/**
 * 
 * @export
 * @interface ModelFileImageUploadSuccessResultModel
 */
export interface ModelFileImageUploadSuccessResultModel {
    /**
     * 
     * @type {string}
     * @memberof ModelFileImageUploadSuccessResultModel
     */
    url?: string;
}
/**
 * 
 * @export
 * @interface ModelLogin
 */
export interface ModelLogin {
    /**
     * 
     * @type {string}
     * @memberof ModelLogin
     */
    password: string;
    /**
     * 
     * @type {string}
     * @memberof ModelLogin
     */
    userID: string;
}
/**
 * 
 * @export
 * @interface ModelLoginSuccessMessage
 */
export interface ModelLoginSuccessMessage {
    /**
     * 
     * @type {string}
     * @memberof ModelLoginSuccessMessage
     */
    expiredAt?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelLoginSuccessMessage
     */
    token?: string;
}
/**
 * 
 * @export
 * @interface ModelPost
 */
export interface ModelPost {
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    body?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    plainBody?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    published?: boolean;
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelPost
     */
    tags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    thumbnail?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    updatedAt?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    userID?: string;
}
/**
 * 
 * @export
 * @interface ModelPostCreateBody
 */
export interface ModelPostCreateBody {
    /**
     * 
     * @type {string}
     * @memberof ModelPostCreateBody
     */
    body: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPostCreateBody
     */
    plainBody: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelPostCreateBody
     */
    tags: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof ModelPostCreateBody
     */
    thumbnail: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPostCreateBody
     */
    title: string;
}
/**
 * 
 * @export
 * @interface ModelPostCreateResult
 */
export interface ModelPostCreateResult {
    /**
     * 
     * @type {string}
     * @memberof ModelPostCreateResult
     */
    postID?: string;
}
/**
 * 
 * @export
 * @interface ModelPostDeleteResult
 */
export interface ModelPostDeleteResult {
    /**
     * 
     * @type {string}
     * @memberof ModelPostDeleteResult
     */
    postID?: string;
}
/**
 * 
 * @export
 * @interface ModelPostListModel
 */
export interface ModelPostListModel {
    /**
     * 
     * @type {number}
     * @memberof ModelPostListModel
     */
    perPageCount?: number;
    /**
     * 
     * @type {Array<ModelPost>}
     * @memberof ModelPostListModel
     */
    posts?: Array<ModelPost>;
    /**
     * 
     * @type {number}
     * @memberof ModelPostListModel
     */
    totalCount?: number;
}
/**
 * 
 * @export
 * @interface ModelTagCreateFailResult
 */
export interface ModelTagCreateFailResult {
    /**
     * 
     * @type {string}
     * @memberof ModelTagCreateFailResult
     */
    message?: string;
}
/**
 * 
 * @export
 * @interface ModelTagCreatedResult
 */
export interface ModelTagCreatedResult {
    /**
     * 
     * @type {string}
     * @memberof ModelTagCreatedResult
     */
    tagName?: string;
}
/**
 * 
 * @export
 * @interface ModelUnauthorizedMessage
 */
export interface ModelUnauthorizedMessage {
    /**
     * 
     * @type {number}
     * @memberof ModelUnauthorizedMessage
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelUnauthorizedMessage
     */
    message?: string;
}
/**
 * 
 * @export
 * @interface ModelUser
 */
export interface ModelUser {
    /**
     * 
     * @type {string}
     * @memberof ModelUser
     */
    userID?: string;
}

/**
 * AuthApi - axios parameter creator
 * @export
 */
export const AuthApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Login
         * @param {ModelLogin} login Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        authLoginPost: async (login: ModelLogin, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'login' is not null or undefined
            if (login === null || login === undefined) {
                throw new RequiredError('login','Required parameter login was null or undefined when calling authLoginPost.');
            }
            const localVarPath = `/auth/login`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const needsSerialization = (typeof login !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(login !== undefined ? login : {}) : (login || "");

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        authUserGet: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/auth/user`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
                    ? await configuration.apiKey("Authorization")
                    : await configuration.apiKey;
                localVarHeaderParameter["Authorization"] = localVarApiKeyValue;
            }


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * AuthApi - functional programming interface
 * @export
 */
export const AuthApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Login
         * @param {ModelLogin} login Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async authLoginPost(login: ModelLogin, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelLoginSuccessMessage>> {
            const localVarAxiosArgs = await AuthApiAxiosParamCreator(configuration).authLoginPost(login, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async authUserGet(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelUser>> {
            const localVarAxiosArgs = await AuthApiAxiosParamCreator(configuration).authUserGet(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * AuthApi - factory interface
 * @export
 */
export const AuthApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Login
         * @param {ModelLogin} login Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        authLoginPost(login: ModelLogin, options?: any): AxiosPromise<ModelLoginSuccessMessage> {
            return AuthApiFp(configuration).authLoginPost(login, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Login
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        authUserGet(options?: any): AxiosPromise<ModelUser> {
            return AuthApiFp(configuration).authUserGet(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * AuthApi - object-oriented interface
 * @export
 * @class AuthApi
 * @extends {BaseAPI}
 */
export class AuthApi extends BaseAPI {
    /**
     * 
     * @summary Login
     * @param {ModelLogin} login Login
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public authLoginPost(login: ModelLogin, options?: any) {
        return AuthApiFp(this.configuration).authLoginPost(login, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Login
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AuthApi
     */
    public authUserGet(options?: any) {
        return AuthApiFp(this.configuration).authUserGet(options).then((request) => request(this.axios, this.basePath));
    }

}


/**
 * ContentsApi - axios parameter creator
 * @export
 */
export const ContentsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Upload a image file of base64 encoded data url type
         * @param {ModelFileImageUploadModel} message Image Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        contentsPost: async (message: ModelFileImageUploadModel, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'message' is not null or undefined
            if (message === null || message === undefined) {
                throw new RequiredError('message','Required parameter message was null or undefined when calling contentsPost.');
            }
            const localVarPath = `/contents`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const needsSerialization = (typeof message !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(message !== undefined ? message : {}) : (message || "");

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * ContentsApi - functional programming interface
 * @export
 */
export const ContentsApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Upload a image file of base64 encoded data url type
         * @param {ModelFileImageUploadModel} message Image Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async contentsPost(message: ModelFileImageUploadModel, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelFileImageUploadSuccessResultModel>> {
            const localVarAxiosArgs = await ContentsApiAxiosParamCreator(configuration).contentsPost(message, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * ContentsApi - factory interface
 * @export
 */
export const ContentsApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Upload a image file of base64 encoded data url type
         * @param {ModelFileImageUploadModel} message Image Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        contentsPost(message: ModelFileImageUploadModel, options?: any): AxiosPromise<ModelFileImageUploadSuccessResultModel> {
            return ContentsApiFp(configuration).contentsPost(message, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * ContentsApi - object-oriented interface
 * @export
 * @class ContentsApi
 * @extends {BaseAPI}
 */
export class ContentsApi extends BaseAPI {
    /**
     * 
     * @summary Upload a image file of base64 encoded data url type
     * @param {ModelFileImageUploadModel} message Image Data
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ContentsApi
     */
    public contentsPost(message: ModelFileImageUploadModel, options?: any) {
        return ContentsApiFp(this.configuration).contentsPost(message, options).then((request) => request(this.axios, this.basePath));
    }

}


/**
 * PostsApi - axios parameter creator
 * @export
 */
export const PostsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get posts with pagination
         * @param {number} [page] Page
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsGet: async (page?: number, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/posts`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (page !== undefined) {
                localVarQueryParameter['page'] = page;
            }


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Delete single post
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsIdDelete: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling postsIdDelete.');
            }
            const localVarPath = `/posts/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
                    ? await configuration.apiKey("Authorization")
                    : await configuration.apiKey;
                localVarHeaderParameter["Authorization"] = localVarApiKeyValue;
            }


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get single post with specific id
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsIdGet: async (id: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling postsIdGet.');
            }
            const localVarPath = `/posts/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create single post
         * @param {ModelPostCreateBody} message Post Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsPost: async (message: ModelPostCreateBody, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'message' is not null or undefined
            if (message === null || message === undefined) {
                throw new RequiredError('message','Required parameter message was null or undefined when calling postsPost.');
            }
            const localVarPath = `/posts`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            if (configuration && configuration.apiKey) {
                const localVarApiKeyValue = typeof configuration.apiKey === 'function'
                    ? await configuration.apiKey("Authorization")
                    : await configuration.apiKey;
                localVarHeaderParameter["Authorization"] = localVarApiKeyValue;
            }


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const needsSerialization = (typeof message !== "string") || localVarRequestOptions.headers['Content-Type'] === 'application/json';
            localVarRequestOptions.data =  needsSerialization ? JSON.stringify(message !== undefined ? message : {}) : (message || "");

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PostsApi - functional programming interface
 * @export
 */
export const PostsApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get posts with pagination
         * @param {number} [page] Page
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async postsGet(page?: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPostListModel>> {
            const localVarAxiosArgs = await PostsApiAxiosParamCreator(configuration).postsGet(page, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Delete single post
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async postsIdDelete(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPostDeleteResult>> {
            const localVarAxiosArgs = await PostsApiAxiosParamCreator(configuration).postsIdDelete(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Get single post with specific id
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async postsIdGet(id: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPost>> {
            const localVarAxiosArgs = await PostsApiAxiosParamCreator(configuration).postsIdGet(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Create single post
         * @param {ModelPostCreateBody} message Post Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async postsPost(message: ModelPostCreateBody, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelPostCreateResult>> {
            const localVarAxiosArgs = await PostsApiAxiosParamCreator(configuration).postsPost(message, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * PostsApi - factory interface
 * @export
 */
export const PostsApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Get posts with pagination
         * @param {number} [page] Page
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsGet(page?: number, options?: any): AxiosPromise<ModelPostListModel> {
            return PostsApiFp(configuration).postsGet(page, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Delete single post
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsIdDelete(id: string, options?: any): AxiosPromise<ModelPostDeleteResult> {
            return PostsApiFp(configuration).postsIdDelete(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get single post with specific id
         * @param {string} id Post ID
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsIdGet(id: string, options?: any): AxiosPromise<ModelPost> {
            return PostsApiFp(configuration).postsIdGet(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create single post
         * @param {ModelPostCreateBody} message Post Data
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsPost(message: ModelPostCreateBody, options?: any): AxiosPromise<ModelPostCreateResult> {
            return PostsApiFp(configuration).postsPost(message, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * PostsApi - object-oriented interface
 * @export
 * @class PostsApi
 * @extends {BaseAPI}
 */
export class PostsApi extends BaseAPI {
    /**
     * 
     * @summary Get posts with pagination
     * @param {number} [page] Page
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostsApi
     */
    public postsGet(page?: number, options?: any) {
        return PostsApiFp(this.configuration).postsGet(page, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Delete single post
     * @param {string} id Post ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostsApi
     */
    public postsIdDelete(id: string, options?: any) {
        return PostsApiFp(this.configuration).postsIdDelete(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get single post with specific id
     * @param {string} id Post ID
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostsApi
     */
    public postsIdGet(id: string, options?: any) {
        return PostsApiFp(this.configuration).postsIdGet(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create single post
     * @param {ModelPostCreateBody} message Post Data
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostsApi
     */
    public postsPost(message: ModelPostCreateBody, options?: any) {
        return PostsApiFp(this.configuration).postsPost(message, options).then((request) => request(this.axios, this.basePath));
    }

}


/**
 * TagsApi - axios parameter creator
 * @export
 */
export const TagsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * Get all tags
         * @summary Get tags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tagsGet: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/tags`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * Create a new tag
         * @summary Create tag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tagsPost: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/tags`;
            const localVarUrlObj = globalImportUrl.parse(localVarPath, true);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }
            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarUrlObj.query = {...localVarUrlObj.query, ...localVarQueryParameter, ...options.query};
            // fix override query string Detail: https://stackoverflow.com/a/7517673/1077943
            delete localVarUrlObj.search;
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: globalImportUrl.format(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * TagsApi - functional programming interface
 * @export
 */
export const TagsApiFp = function(configuration?: Configuration) {
    return {
        /**
         * Get all tags
         * @summary Get tags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async tagsGet(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<string>>> {
            const localVarAxiosArgs = await TagsApiAxiosParamCreator(configuration).tagsGet(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * Create a new tag
         * @summary Create tag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async tagsPost(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ModelTagCreatedResult>> {
            const localVarAxiosArgs = await TagsApiAxiosParamCreator(configuration).tagsPost(options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: basePath + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * TagsApi - factory interface
 * @export
 */
export const TagsApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * Get all tags
         * @summary Get tags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tagsGet(options?: any): AxiosPromise<Array<string>> {
            return TagsApiFp(configuration).tagsGet(options).then((request) => request(axios, basePath));
        },
        /**
         * Create a new tag
         * @summary Create tag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        tagsPost(options?: any): AxiosPromise<ModelTagCreatedResult> {
            return TagsApiFp(configuration).tagsPost(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * TagsApi - object-oriented interface
 * @export
 * @class TagsApi
 * @extends {BaseAPI}
 */
export class TagsApi extends BaseAPI {
    /**
     * Get all tags
     * @summary Get tags
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TagsApi
     */
    public tagsGet(options?: any) {
        return TagsApiFp(this.configuration).tagsGet(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * Create a new tag
     * @summary Create tag
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof TagsApi
     */
    public tagsPost(options?: any) {
        return TagsApiFp(this.configuration).tagsPost(options).then((request) => request(this.axios, this.basePath));
    }

}


