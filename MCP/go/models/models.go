package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ImageDescription represents the ImageDescription schema from the OpenAPI specification
type ImageDescription struct {
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
	Description ImageDescriptionDetails `json:"description,omitempty"` // A collection of content tags, along with a list of captions sorted by confidence level, and image metadata.
}

// CategoryDetail represents the CategoryDetail schema from the OpenAPI specification
type CategoryDetail struct {
	Celebrities []CelebritiesModel `json:"celebrities,omitempty"` // An array of celebrities if any identified.
	Landmarks []LandmarksModel `json:"landmarks,omitempty"` // An array of landmarks if any identified.
}

// LandmarksModel represents the LandmarksModel schema from the OpenAPI specification
type LandmarksModel struct {
	Confidence float64 `json:"confidence,omitempty"` // Confidence level for the landmark recognition as a value ranging from 0 to 1.
	Name string `json:"name,omitempty"` // Name of the landmark.
}

// ColorInfo represents the ColorInfo schema from the OpenAPI specification
type ColorInfo struct {
	Dominantcolorbackground string `json:"dominantColorBackground,omitempty"` // Possible dominant background color.
	Dominantcolorforeground string `json:"dominantColorForeground,omitempty"` // Possible dominant foreground color.
	Dominantcolors []string `json:"dominantColors,omitempty"` // An array of possible dominant colors.
	Isbwimg bool `json:"isBWImg,omitempty"` // A value indicating if the image is black and white.
	Accentcolor string `json:"accentColor,omitempty"` // Possible accent color.
}

// ImageTag represents the ImageTag schema from the OpenAPI specification
type ImageTag struct {
	Hint string `json:"hint,omitempty"` // Optional hint/details for this tag.
	Name string `json:"name,omitempty"` // Name of the entity.
	Confidence float64 `json:"confidence,omitempty"` // The level of confidence that the entity was observed.
}

// ImageAnalysis represents the ImageAnalysis schema from the OpenAPI specification
type ImageAnalysis struct {
	Adult AdultInfo `json:"adult,omitempty"` // An object describing whether the image contains adult-oriented content and/or is racy.
	Categories []Category `json:"categories,omitempty"` // An array indicating identified categories.
	Color ColorInfo `json:"color,omitempty"` // An object providing additional metadata describing color attributes.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
	Imagetype ImageType `json:"imageType,omitempty"` // An object providing possible image types and matching confidence levels.
	Tags []ImageTag `json:"tags,omitempty"` // A list of tags with confidence level.
	Brands []DetectedBrand `json:"brands,omitempty"` // Array of brands detected in the image.
	Description ImageDescriptionDetails `json:"description,omitempty"` // A collection of content tags, along with a list of captions sorted by confidence level, and image metadata.
	Objects []DetectedObject `json:"objects,omitempty"` // Array of objects describing what was detected in the image.
	Faces []FaceDescription `json:"faces,omitempty"` // An array of possible faces within the image.
}

// FaceRectangle represents the FaceRectangle schema from the OpenAPI specification
type FaceRectangle struct {
	Height int `json:"height,omitempty"` // Height measured from the top-left point of the face, in pixels.
	Left int `json:"left,omitempty"` // X-coordinate of the top left point of the face, in pixels.
	Top int `json:"top,omitempty"` // Y-coordinate of the top left point of the face, in pixels.
	Width int `json:"width,omitempty"` // Width measured from the top-left point of the face, in pixels.
}

// DomainModelResults represents the DomainModelResults schema from the OpenAPI specification
type DomainModelResults struct {
	Result map[string]interface{} `json:"result,omitempty"` // Model-specific response.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
}

// OcrLine represents the OcrLine schema from the OpenAPI specification
type OcrLine struct {
	Boundingbox string `json:"boundingBox,omitempty"` // Bounding box of a recognized line. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
	Words []OcrWord `json:"words,omitempty"` // An array of objects, where each object represents a recognized word.
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Detail CategoryDetail `json:"detail,omitempty"` // An object describing additional category details.
	Name string `json:"name,omitempty"` // Name of the category.
	Score float64 `json:"score,omitempty"` // Scoring of the category.
}

// ObjectHierarchy represents the ObjectHierarchy schema from the OpenAPI specification
type ObjectHierarchy struct {
	Confidence float64 `json:"confidence,omitempty"` // Confidence score of having observed the object in the image, as a value ranging from 0 to 1.
	Object string `json:"object,omitempty"` // Label for the object.
	Parent ObjectHierarchy `json:"parent,omitempty"` // An object detected inside an image.
}

// CelebrityResults represents the CelebrityResults schema from the OpenAPI specification
type CelebrityResults struct {
	Celebrities []CelebritiesModel `json:"celebrities,omitempty"` // List of celebrities recognized in the image.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
}

// ImageMetadata represents the ImageMetadata schema from the OpenAPI specification
type ImageMetadata struct {
	Format string `json:"format,omitempty"` // Image format.
	Height int `json:"height,omitempty"` // Image height, in pixels.
	Width int `json:"width,omitempty"` // Image width, in pixels.
}

// ListModelsResult represents the ListModelsResult schema from the OpenAPI specification
type ListModelsResult struct {
	Models []ModelDescription `json:"models,omitempty"` // An array of supported models.
}

// ComputerVisionError represents the ComputerVisionError schema from the OpenAPI specification
type ComputerVisionError struct {
	Code interface{} `json:"code"` // The error code.
	Message string `json:"message"` // A message explaining the error reported by the service.
	Requestid string `json:"requestId,omitempty"` // A unique request identifier.
}

// DetectedBrand represents the DetectedBrand schema from the OpenAPI specification
type DetectedBrand struct {
	Confidence float64 `json:"confidence,omitempty"` // Confidence score of having observed the brand in the image, as a value ranging from 0 to 1.
	Name string `json:"name,omitempty"` // Label for the brand.
	Rectangle BoundingRect `json:"rectangle,omitempty"` // A bounding box for an area inside an image.
}

// OcrResult represents the OcrResult schema from the OpenAPI specification
type OcrResult struct {
	Language string `json:"language,omitempty"` // The BCP-47 language code of the text in the image.
	Orientation string `json:"orientation,omitempty"` // Orientation of the text recognized in the image. The value (up, down, left, or right) refers to the direction that the top of the recognized text is facing, after the image has been rotated around its center according to the detected text angle (see textAngle property).
	Regions []OcrRegion `json:"regions,omitempty"` // An array of objects, where each object represents a region of recognized text.
	Textangle float64 `json:"textAngle,omitempty"` // The angle, in degrees, of the detected text with respect to the closest horizontal or vertical direction. After rotating the input image clockwise by this angle, the recognized text lines become horizontal or vertical. In combination with the orientation property it can be used to overlay recognition results correctly on the original image, by rotating either the original image or recognition results by a suitable angle around the center of the original image. If the angle cannot be confidently detected, this property is not present. If the image contains text at different angles, only part of the text will be recognized correctly.
}

// ImageType represents the ImageType schema from the OpenAPI specification
type ImageType struct {
	Linedrawingtype int `json:"lineDrawingType,omitempty"` // Confidence level that the image is a line drawing.
	Cliparttype int `json:"clipArtType,omitempty"` // Confidence level that the image is a clip art.
}

// BoundingRect represents the BoundingRect schema from the OpenAPI specification
type BoundingRect struct {
	H int `json:"h,omitempty"` // Height measured from the top-left point of the area, in pixels.
	W int `json:"w,omitempty"` // Width measured from the top-left point of the area, in pixels.
	X int `json:"x,omitempty"` // X-coordinate of the top left point of the area, in pixels.
	Y int `json:"y,omitempty"` // Y-coordinate of the top left point of the area, in pixels.
}

// ImageDescriptionDetails represents the ImageDescriptionDetails schema from the OpenAPI specification
type ImageDescriptionDetails struct {
	Captions []ImageCaption `json:"captions,omitempty"` // A list of captions, sorted by confidence level.
	Tags []string `json:"tags,omitempty"` // A collection of image tags.
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"url"` // Publicly reachable URL of an image.
}

// OcrWord represents the OcrWord schema from the OpenAPI specification
type OcrWord struct {
	Boundingbox string `json:"boundingBox,omitempty"` // Bounding box of a recognized word. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
	Text string `json:"text,omitempty"` // String value of a recognized word.
}

// TagResult represents the TagResult schema from the OpenAPI specification
type TagResult struct {
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
	Tags []ImageTag `json:"tags,omitempty"` // A list of tags with confidence level.
}

// ImageCaption represents the ImageCaption schema from the OpenAPI specification
type ImageCaption struct {
	Confidence float64 `json:"confidence,omitempty"` // The level of confidence the service has in the caption.
	Text string `json:"text,omitempty"` // The text of the caption.
}

// DetectResult represents the DetectResult schema from the OpenAPI specification
type DetectResult struct {
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Objects []DetectedObject `json:"objects,omitempty"` // An array of detected objects.
}

// AreaOfInterestResult represents the AreaOfInterestResult schema from the OpenAPI specification
type AreaOfInterestResult struct {
	Areaofinterest BoundingRect `json:"areaOfInterest,omitempty"` // A bounding box for an area inside an image.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
}

// LandmarkResults represents the LandmarkResults schema from the OpenAPI specification
type LandmarkResults struct {
	Landmarks []LandmarksModel `json:"landmarks,omitempty"` // List of landmarks recognized in the image.
	Metadata ImageMetadata `json:"metadata,omitempty"` // Image metadata.
	Requestid string `json:"requestId,omitempty"` // Id of the REST API request.
}

// OcrRegion represents the OcrRegion schema from the OpenAPI specification
type OcrRegion struct {
	Lines []OcrLine `json:"lines,omitempty"` // An array of recognized lines of text.
	Boundingbox string `json:"boundingBox,omitempty"` // Bounding box of a recognized region. The four integers represent the x-coordinate of the left edge, the y-coordinate of the top edge, width, and height of the bounding box, in the coordinate system of the input image, after it has been rotated around its center according to the detected text angle (see textAngle property), with the origin at the top-left corner, and the y-axis pointing down.
}

// AdultInfo represents the AdultInfo schema from the OpenAPI specification
type AdultInfo struct {
	Adultscore float64 `json:"adultScore,omitempty"` // Score from 0 to 1 that indicates how much the content is considered adult-oriented within the image.
	Isadultcontent bool `json:"isAdultContent,omitempty"` // A value indicating if the image contains adult-oriented content.
	Isracycontent bool `json:"isRacyContent,omitempty"` // A value indicating if the image is racy.
	Racyscore float64 `json:"racyScore,omitempty"` // Score from 0 to 1 that indicates how suggestive is the image.
}

// ModelDescription represents the ModelDescription schema from the OpenAPI specification
type ModelDescription struct {
	Name string `json:"name,omitempty"` // The name of the model.
	Categories []string `json:"categories,omitempty"` // Categories of the model.
}

// DetectedObject represents the DetectedObject schema from the OpenAPI specification
type DetectedObject struct {
	Confidence float64 `json:"confidence,omitempty"` // Confidence score of having observed the object in the image, as a value ranging from 0 to 1.
	Object string `json:"object,omitempty"` // Label for the object.
	Parent ObjectHierarchy `json:"parent,omitempty"` // An object detected inside an image.
	Rectangle BoundingRect `json:"rectangle,omitempty"` // A bounding box for an area inside an image.
}

// FaceDescription represents the FaceDescription schema from the OpenAPI specification
type FaceDescription struct {
	Age int `json:"age,omitempty"` // Possible age of the face.
	Facerectangle FaceRectangle `json:"faceRectangle,omitempty"` // An object describing face rectangle.
	Gender string `json:"gender,omitempty"` // Possible gender of the face.
}

// CelebritiesModel represents the CelebritiesModel schema from the OpenAPI specification
type CelebritiesModel struct {
	Facerectangle FaceRectangle `json:"faceRectangle,omitempty"` // An object describing face rectangle.
	Name string `json:"name,omitempty"` // Name of the celebrity.
	Confidence float64 `json:"confidence,omitempty"` // Confidence level for the celebrity recognition as a value ranging from 0 to 1.
}
