function createHeader(headerType) {
  let contentType;
  if (headerType === "file") {
    contentType = "multipart/form-data";
  } else if (headerType === "json") {
    contentType = "application/json";
  } else if (headerType === "data") {
    contentType = "multipart/form-data";
  } else {
    contentType = "application/x-www-form-urlencoded";
  }

  return {
    "Content-Type": contentType,
  };
}

export { createHeader };
