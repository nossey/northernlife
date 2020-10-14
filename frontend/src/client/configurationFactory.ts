import  {Configuration} from "./configuration"

export function buildConfiguration(): Configuration {
  if (process.client)
    return new Configuration({basePath: "api/"});
  if (process.server){
    const basePath = (process.env.API_BASE_URL) ? `${process.env.API_BASE_URL}/api/v1` : "http://localhost:9000/api/v1";
    console.log(basePath);
    return new Configuration(({basePath}));
  }
  throw new Error();
};
