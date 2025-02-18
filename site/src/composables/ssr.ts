export interface Image {
  imageUlid: string; // db:"image_ulid"
  image: string; // db:"image"
  originalHeight: number; // db:"original_height"
  originalWidth: number; // db:"original_width"
  top: number; // db:"top"
  left: number; // db:"left"
  scale: number; // db:"scale"
  cropHeight: number; // db:"crop_height"
  cropWidth: number; // db:"crop_width"
  siteUlid: string; // db:"site_ulid"
  dateCreated: string; // db:"date_created"
  dateModified: string; // db:"date_modified"
  isProcessed: boolean; // db:"is_processed"
}

interface PageWithBlockWithImage {
  pageUlid?: string;
  title?: string;
  summary?: string;
  dateCreated?: string;
  dateModified?: string;
  siteUlid?: string;
  slug?: string;
  blocks?: BlockWithImage[];
}

interface BlockWithImage {
  blockUlid?: string;
  title?: string;
  kind?: string;
  contentOneHtml?: string;
  contentTwoHtml?: string;
  contentThreeHtml?: string;
  contentFourHtml?: string;
  imageOne?: Image;
  imageTwo?: Image;
  imageThree?: Image;
  imageFour?: Image;
  dateCreated?: string;
  dateModified?: string;
  pageUlid?: string;
  siteUlid?: string;
  sortPosition?: number; // Use number for int32
}

interface NavItem {
  slug: string;
  title: string;
}

interface Metadata {
  mainNav: NavItem[];
}

interface Payload {
  page: PageWithBlockWithImage;
  metadata: Metadata;
}

export function ssrFetch(url: string): Promise<PageWithBlockWithImage> {
  return new Promise((resolve, reject) => {
    if (typeof window === "undefined") {
      if ((globalThis as any).payload) {
        let payload: Payload = (globalThis as any).payload;
        if (typeof payload === "string") {
          let pl: Payload = JSON.parse(payload);
          resolve(pl.page);
        }
      }
    } else {
      fetch("/page" + url)
        .then((resp) => {
          return resp.json();
        })
        .then((data) => {
          resolve(data);
        });
    }
  });
}

export function ssrMetadata(): Promise<any> {
  return new Promise((resolve, reject) => {
    try {
      fetch("/metadata")
        .then((resp) => {
          return resp.json();
        })
        .then((data) => {
          resolve(data);
        });
    } catch (err: any) {
      resolve({
        ssrError: err,
      });
    }
  });
}

export function useSSRMetadata() {
  let meta = ref<Metadata | undefined>();
  onMounted(async () => {
    meta.value = await ssrMetadata();
  });

  if ((globalThis as any).payload) {
    let metaVal;
    let payload = (globalThis as any).payload;
    try {
      if (typeof payload === "string") {
        metaVal = JSON.parse(payload).metadata;
      }
    } catch (err: any) {
      // HANDLE THE ERROR
      metaVal({
        ssrError: err,
      });
    }
    meta.value = metaVal
  }
  return meta;
}
