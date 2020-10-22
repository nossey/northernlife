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
 * @interface ModelLoginFailMessage
 */
export interface ModelLoginFailMessage {
    /**
     * 
     * @type {number}
     * @memberof ModelLoginFailMessage
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelLoginFailMessage
     */
    message?: string;
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
    expired_at?: string;
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
    created_at?: string;
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
    plain_body?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelPost
     */
    published?: boolean;
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
    updated_at?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelPost
     */
    user_id?: string;
}
/**
 * 
 * @export
 * @interface ModelPostListModel
 */
export interface ModelPostListModel {
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
    total_count?: number;
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

}


/**
 * PostsApi - axios parameter creator
 * @export
 */
export const PostsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Retrive posts
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsGet: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/posts`;
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
         * @summary Retrive posts
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async postsGet(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelPostListModel>>> {
            const localVarAxiosArgs = await PostsApiAxiosParamCreator(configuration).postsGet(options);
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
         * @summary Retrive posts
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        postsGet(options?: any): AxiosPromise<Array<ModelPostListModel>> {
            return PostsApiFp(configuration).postsGet(options).then((request) => request(axios, basePath));
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
     * @summary Retrive posts
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PostsApi
     */
    public postsGet(options?: any) {
        return PostsApiFp(this.configuration).postsGet(options).then((request) => request(this.axios, this.basePath));
    }

}


