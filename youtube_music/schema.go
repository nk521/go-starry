package youtube_music

type SchemaGetHomePage struct {
	ResponseContext    ResponseContext `json:"responseContext"`
	Contents           Contents        `json:"contents"`
	TrackingParams     string          `json:"trackingParams"`
	MaxAgeStoreSeconds int64           `json:"maxAgeStoreSeconds"`
	Background         ThumbnailClass  `json:"background"`
}

type ThumbnailClass struct {
	MusicThumbnailRenderer BackgroundMusicThumbnailRenderer `json:"musicThumbnailRenderer"`
}

type BackgroundMusicThumbnailRenderer struct {
	Thumbnail      MusicThumbnailRendererThumbnail `json:"thumbnail"`
	ThumbnailCrop  ThumbnailCrop                   `json:"thumbnailCrop"`
	ThumbnailScale ThumbnailScale                  `json:"thumbnailScale"`
	TrackingParams string                          `json:"trackingParams"`
}

type MusicThumbnailRendererThumbnail struct {
	Thumbnails []ThumbnailElement `json:"thumbnails"`
}

type ThumbnailElement struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type Contents struct {
	SingleColumnBrowseResultsRenderer SingleColumnBrowseResultsRenderer `json:"singleColumnBrowseResultsRenderer"`
}

type SingleColumnBrowseResultsRenderer struct {
	Tabs []Tab `json:"tabs"`
}

type Tab struct {
	TabRenderer TabRenderer `json:"tabRenderer"`
}

type TabRenderer struct {
	Endpoint       Endpoint           `json:"endpoint"`
	Title          string             `json:"title"`
	Selected       bool               `json:"selected"`
	Content        TabRendererContent `json:"content"`
	Icon           Icon               `json:"icon"`
	TabIdentifier  TabIdentifier      `json:"tabIdentifier"`
	TrackingParams string             `json:"trackingParams"`
}

type TabRendererContent struct {
	SectionListRenderer SectionListRenderer `json:"sectionListRenderer"`
}

type SectionListRenderer struct {
	Contents       []SectionListRendererContent `json:"contents"`
	Continuations  []Continuation               `json:"continuations"`
	TrackingParams string                       `json:"trackingParams"`
	Header         SectionListRendererHeader    `json:"header"`
}

type SectionListRendererContent struct {
	MusicCarouselShelfRenderer    MusicCarouselShelfRenderer    `json:"musicCarouselShelfRenderer"`
	MusicDescriptionShelfRenderer MusicDescriptionShelfRenderer `json:"musicDescriptionShelfRenderer"`
}

type MusicCarouselShelfRenderer struct {
	Header            MusicCarouselShelfRendererHeader    `json:"header"`
	Contents          []MusicCarouselShelfRendererContent `json:"contents"`
	TrackingParams    string                              `json:"trackingParams"`
	ItemSize          string                              `json:"itemSize"`
	NumItemsPerColumn *string                             `json:"numItemsPerColumn,omitempty"`
}

type MusicCarouselShelfRendererContent struct {
	MusicTwoRowItemRenderer         *MusicTwoRowItemRenderer         `json:"musicTwoRowItemRenderer,omitempty"`
	MusicResponsiveListItemRenderer *MusicResponsiveListItemRenderer `json:"musicResponsiveListItemRenderer,omitempty"`
}

type MusicDescriptionShelfRenderer struct {
	Header      MusicDescriptionShelfRendererHeaderOrDescription `json:"header,omitempty"`
	Description MusicDescriptionShelfRendererHeaderOrDescription `json:"description,omitempty"`
}

type MusicDescriptionShelfRendererHeaderOrDescription struct {
	Runs []MusicDescriptionShelfRendererHeaderRuns `json:"runs,omitempty"`
}

type MusicDescriptionShelfRendererHeaderRuns struct {
	Text string `json:"text,omitempty"`
}

type MusicResponsiveListItemRenderer struct {
	TrackingParams         string                              `json:"trackingParams"`
	Thumbnail              ThumbnailClass                      `json:"thumbnail"`
	Overlay                Overlay                             `json:"overlay"`
	FlexColumns            []FlexColumn                        `json:"flexColumns"`
	Menu                   MusicResponsiveListItemRendererMenu `json:"menu"`
	PlaylistItemData       PlaylistItemData                    `json:"playlistItemData"`
	FlexColumnDisplayStyle string                              `json:"flexColumnDisplayStyle"`
	ItemHeight             ItemHeight                          `json:"itemHeight"`
	Badges                 []Badge                             `json:"badges,omitempty"`
}

type Badge struct {
	MusicInlineBadgeRenderer MusicInlineBadgeRenderer `json:"musicInlineBadgeRenderer"`
}

type MusicInlineBadgeRenderer struct {
	TrackingParams    string        `json:"trackingParams"`
	Icon              Icon          `json:"icon"`
	AccessibilityData Accessibility `json:"accessibilityData"`
}

type Accessibility struct {
	AccessibilityData AccessibilityData `json:"accessibilityData"`
}

type AccessibilityData struct {
	Label string `json:"label"`
}

type Icon struct {
	IconType IconType `json:"iconType"`
}

type FlexColumn struct {
	MusicResponsiveListItemFlexColumnRenderer MusicResponsiveListItemFlexColumnRenderer `json:"musicResponsiveListItemFlexColumnRenderer"`
}

type MusicResponsiveListItemFlexColumnRenderer struct {
	Text            Text            `json:"text"`
	DisplayPriority DisplayPriority `json:"displayPriority"`
}

type Text struct {
	Runs []PurpleRun `json:"runs"`
}

type PurpleRun struct {
	Text               string                    `json:"text"`
	NavigationEndpoint *PurpleNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type PurpleNavigationEndpoint struct {
	ClickTrackingParams string               `json:"clickTrackingParams"`
	WatchEndpoint       *WatchEndpoint       `json:"watchEndpoint,omitempty"`
	BrowseEndpoint      *OnTapBrowseEndpoint `json:"browseEndpoint,omitempty"`
}

type OnTapBrowseEndpoint struct {
	BrowseID                              string                                `json:"browseId"`
	BrowseEndpointContextSupportedConfigs BrowseEndpointContextSupportedConfigs `json:"browseEndpointContextSupportedConfigs"`
}

type BrowseEndpointContextSupportedConfigs struct {
	BrowseEndpointContextMusicConfig BrowseEndpointContextMusicConfig `json:"browseEndpointContextMusicConfig"`
}

type BrowseEndpointContextMusicConfig struct {
	PageType PageType `json:"pageType"`
}

type WatchEndpoint struct {
	VideoID                            string                             `json:"videoId"`
	PlaylistID                         string                             `json:"playlistId"`
	Params                             *Params                            `json:"params,omitempty"`
	LoggingContext                     LoggingContext                     `json:"loggingContext"`
	WatchEndpointMusicSupportedConfigs WatchEndpointMusicSupportedConfigs `json:"watchEndpointMusicSupportedConfigs"`
}

type LoggingContext struct {
	VssLoggingContext VssLoggingContext `json:"vssLoggingContext"`
}

type VssLoggingContext struct {
	SerializedContextData string `json:"serializedContextData"`
}

type WatchEndpointMusicSupportedConfigs struct {
	WatchEndpointMusicConfig WatchEndpointMusicConfig `json:"watchEndpointMusicConfig"`
}

type WatchEndpointMusicConfig struct {
	MusicVideoType MusicVideoType `json:"musicVideoType"`
}

type MusicResponsiveListItemRendererMenu struct {
	MenuRenderer PurpleMenuRenderer `json:"menuRenderer"`
}

type PurpleMenuRenderer struct {
	Items           []PurpleItem     `json:"items"`
	TrackingParams  string           `json:"trackingParams"`
	TopLevelButtons []TopLevelButton `json:"topLevelButtons"`
	Accessibility   Accessibility    `json:"accessibility"`
}

type PurpleItem struct {
	MenuNavigationItemRenderer    *MenuItemRenderer                    `json:"menuNavigationItemRenderer,omitempty"`
	MenuServiceItemRenderer       *MenuItemRenderer                    `json:"menuServiceItemRenderer,omitempty"`
	ToggleMenuServiceItemRenderer *PurpleToggleMenuServiceItemRenderer `json:"toggleMenuServiceItemRenderer,omitempty"`
}

type MenuItemRenderer struct {
	Text               Strapline                                     `json:"text"`
	Icon               Icon                                          `json:"icon"`
	NavigationEndpoint *MenuNavigationItemRendererNavigationEndpoint `json:"navigationEndpoint,omitempty"`
	TrackingParams     string                                        `json:"trackingParams"`
	ServiceEndpoint    *MenuNavigationItemRendererServiceEndpoint    `json:"serviceEndpoint,omitempty"`
}

type MenuNavigationItemRendererNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *WatchEndpoint         `json:"watchEndpoint,omitempty"`
	AddToPlaylistEndpoint *AddToPlaylistEndpoint `json:"addToPlaylistEndpoint,omitempty"`
	BrowseEndpoint        *OnTapBrowseEndpoint   `json:"browseEndpoint,omitempty"`
	ShareEntityEndpoint   *ShareEntityEndpoint   `json:"shareEntityEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
	ConfirmDialogEndpoint *ConfirmDialogEndpoint `json:"confirmDialogEndpoint,omitempty"`
}

type AddToPlaylistEndpoint struct {
	VideoID    *string `json:"videoId,omitempty"`
	PlaylistID *string `json:"playlistId,omitempty"`
}

type ConfirmDialogEndpoint struct {
	Content ConfirmDialogEndpointContent `json:"content"`
}

type ConfirmDialogEndpointContent struct {
	ConfirmDialogRenderer ConfirmDialogRenderer `json:"confirmDialogRenderer"`
}

type ConfirmDialogRenderer struct {
	Title          Strapline   `json:"title"`
	TrackingParams string      `json:"trackingParams"`
	DialogMessages []Strapline `json:"dialogMessages"`
	ConfirmButton  Button      `json:"confirmButton"`
	CancelButton   Button      `json:"cancelButton"`
}

type Button struct {
	ButtonRenderer CancelButtonButtonRenderer `json:"buttonRenderer"`
}

type CancelButtonButtonRenderer struct {
	Style          string                 `json:"style"`
	IsDisabled     bool                   `json:"isDisabled"`
	Text           Strapline              `json:"text"`
	TrackingParams string                 `json:"trackingParams"`
	Command        *ButtonRendererCommand `json:"command,omitempty"`
}

type ButtonRendererCommand struct {
	ClickTrackingParams    string                 `json:"clickTrackingParams"`
	CommandExecutorCommand CommandExecutorCommand `json:"commandExecutorCommand"`
}

type CommandExecutorCommand struct {
	Commands []CommandExecutorCommandCommand `json:"commands"`
}

type CommandExecutorCommandCommand struct {
	ClickTrackingParams  string                `json:"clickTrackingParams"`
	PlaylistEditEndpoint *PlaylistEditEndpoint `json:"playlistEditEndpoint,omitempty"`
	LikeEndpoint         *CommandLikeEndpoint  `json:"likeEndpoint,omitempty"`
}

type CommandLikeEndpoint struct {
	Status Status `json:"status"`
	Target Target `json:"target"`
}

type Target struct {
	PlaylistID string `json:"playlistId"`
}

type PlaylistEditEndpoint struct {
	PlaylistID string                       `json:"playlistId"`
	Actions    []PlaylistEditEndpointAction `json:"actions"`
	Params     string                       `json:"params"`
}

type PlaylistEditEndpointAction struct {
	Action          string  `json:"action"`
	PlaylistPrivacy *string `json:"playlistPrivacy,omitempty"`
}

type Strapline struct {
	Runs []StraplineRun `json:"runs"`
}

type StraplineRun struct {
	Text string `json:"text"`
}

type ShareEntityEndpoint struct {
	SerializedShareEntity string         `json:"serializedShareEntity"`
	SharePanelType        SharePanelType `json:"sharePanelType"`
}

type WatchPlaylistEndpoint struct {
	PlaylistID string `json:"playlistId"`
	Params     string `json:"params"`
}

type MenuNavigationItemRendererServiceEndpoint struct {
	ClickTrackingParams string           `json:"clickTrackingParams"`
	QueueAddEndpoint    QueueAddEndpoint `json:"queueAddEndpoint"`
}

type QueueAddEndpoint struct {
	QueueTarget         QueueTarget               `json:"queueTarget"`
	QueueInsertPosition QueueInsertPosition       `json:"queueInsertPosition"`
	Commands            []QueueAddEndpointCommand `json:"commands"`
}

type QueueAddEndpointCommand struct {
	ClickTrackingParams string           `json:"clickTrackingParams"`
	AddToToastAction    AddToToastAction `json:"addToToastAction"`
}

type AddToToastAction struct {
	Item AddToToastActionItem `json:"item"`
}

type AddToToastActionItem struct {
	NotificationTextRenderer NotificationTextRenderer `json:"notificationTextRenderer"`
}

type NotificationTextRenderer struct {
	SuccessResponseText Strapline `json:"successResponseText"`
	TrackingParams      string    `json:"trackingParams"`
}

type QueueTarget struct {
	VideoID      *string      `json:"videoId,omitempty"`
	OnEmptyQueue OnEmptyQueue `json:"onEmptyQueue"`
	PlaylistID   *string      `json:"playlistId,omitempty"`
}

type OnEmptyQueue struct {
	ClickTrackingParams string                `json:"clickTrackingParams"`
	WatchEndpoint       AddToPlaylistEndpoint `json:"watchEndpoint"`
}

type PurpleToggleMenuServiceItemRenderer struct {
	DefaultText            Strapline             `json:"defaultText"`
	DefaultIcon            Icon                  `json:"defaultIcon"`
	DefaultServiceEndpoint PurpleServiceEndpoint `json:"defaultServiceEndpoint"`
	ToggledText            Strapline             `json:"toggledText"`
	ToggledIcon            Icon                  `json:"toggledIcon"`
	ToggledServiceEndpoint PurpleServiceEndpoint `json:"toggledServiceEndpoint"`
	TrackingParams         string                `json:"trackingParams"`
}

type PurpleServiceEndpoint struct {
	ClickTrackingParams string           `json:"clickTrackingParams"`
	FeedbackEndpoint    FeedbackEndpoint `json:"feedbackEndpoint"`
}

type FeedbackEndpoint struct {
	FeedbackToken string `json:"feedbackToken"`
}

type TopLevelButton struct {
	LikeButtonRenderer LikeButtonRenderer `json:"likeButtonRenderer"`
}

type LikeButtonRenderer struct {
	Target           PlaylistItemData         `json:"target"`
	LikeStatus       Status                   `json:"likeStatus"`
	TrackingParams   string                   `json:"trackingParams"`
	LikesAllowed     bool                     `json:"likesAllowed"`
	ServiceEndpoints []ServiceEndpointElement `json:"serviceEndpoints"`
}

type ServiceEndpointElement struct {
	ClickTrackingParams string                      `json:"clickTrackingParams"`
	LikeEndpoint        ServiceEndpointLikeEndpoint `json:"likeEndpoint"`
}

type ServiceEndpointLikeEndpoint struct {
	Status  Status               `json:"status"`
	Target  PlaylistItemData     `json:"target"`
	Actions []LikeEndpointAction `json:"actions,omitempty"`
}

type LikeEndpointAction struct {
	ClickTrackingParams             string                          `json:"clickTrackingParams"`
	MusicLibraryStatusUpdateCommand MusicLibraryStatusUpdateCommand `json:"musicLibraryStatusUpdateCommand"`
}

type MusicLibraryStatusUpdateCommand struct {
	LibraryStatus             LibraryStatus `json:"libraryStatus"`
	AddToLibraryFeedbackToken string        `json:"addToLibraryFeedbackToken"`
}

type PlaylistItemData struct {
	VideoID string `json:"videoId"`
}

type Overlay struct {
	MusicItemThumbnailOverlayRenderer OverlayMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type OverlayMusicItemThumbnailOverlayRenderer struct {
	Background      MusicItemThumbnailOverlayRendererBackground `json:"background"`
	Content         PurpleContent                               `json:"content"`
	ContentPosition ContentPosition                             `json:"contentPosition"`
	DisplayStyle    DisplayStyle                                `json:"displayStyle"`
}

type MusicItemThumbnailOverlayRendererBackground struct {
	VerticalGradient VerticalGradient `json:"verticalGradient"`
}

type VerticalGradient struct {
	GradientLayerColors []string `json:"gradientLayerColors"`
}

type PurpleContent struct {
	MusicPlayButtonRenderer PurpleMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type PurpleMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint PurplePlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      int64                        `json:"activeScaleFactor"`
	ButtonSize             PurpleButtonSize             `json:"buttonSize"`
	RippleTarget           RippleTarget                 `json:"rippleTarget"`
	AccessibilityPlayData  Accessibility                `json:"accessibilityPlayData"`
	AccessibilityPauseData Accessibility                `json:"accessibilityPauseData"`
}

type PurplePlayNavigationEndpoint struct {
	ClickTrackingParams string        `json:"clickTrackingParams"`
	WatchEndpoint       WatchEndpoint `json:"watchEndpoint"`
}

type MusicTwoRowItemRenderer struct {
	ThumbnailRenderer  ThumbnailClass                            `json:"thumbnailRenderer"`
	AspectRatio        AspectRatio                               `json:"aspectRatio"`
	Title              MusicTwoRowItemRendererTitle              `json:"title"`
	Subtitle           Subtitle                                  `json:"subtitle"`
	NavigationEndpoint MusicTwoRowItemRendererNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                                    `json:"trackingParams"`
	Menu               MusicTwoRowItemRendererMenu               `json:"menu"`
	ThumbnailOverlay   ThumbnailOverlay                          `json:"thumbnailOverlay"`
	SubtitleBadges     []Badge                                   `json:"subtitleBadges,omitempty"`
}

type MusicTwoRowItemRendererMenu struct {
	MenuRenderer FluffyMenuRenderer `json:"menuRenderer"`
}

type FluffyMenuRenderer struct {
	Items          []FluffyItem  `json:"items"`
	TrackingParams string        `json:"trackingParams"`
	Accessibility  Accessibility `json:"accessibility"`
}

type FluffyItem struct {
	MenuNavigationItemRenderer    *MenuItemRenderer                    `json:"menuNavigationItemRenderer,omitempty"`
	MenuServiceItemRenderer       *MenuItemRenderer                    `json:"menuServiceItemRenderer,omitempty"`
	ToggleMenuServiceItemRenderer *FluffyToggleMenuServiceItemRenderer `json:"toggleMenuServiceItemRenderer,omitempty"`
}

type FluffyToggleMenuServiceItemRenderer struct {
	DefaultText            Strapline             `json:"defaultText"`
	DefaultIcon            Icon                  `json:"defaultIcon"`
	DefaultServiceEndpoint FluffyServiceEndpoint `json:"defaultServiceEndpoint"`
	ToggledText            Strapline             `json:"toggledText"`
	ToggledIcon            Icon                  `json:"toggledIcon"`
	ToggledServiceEndpoint FluffyServiceEndpoint `json:"toggledServiceEndpoint"`
	TrackingParams         string                `json:"trackingParams"`
}

type FluffyServiceEndpoint struct {
	ClickTrackingParams string                              `json:"clickTrackingParams"`
	FeedbackEndpoint    *FeedbackEndpoint                   `json:"feedbackEndpoint,omitempty"`
	LikeEndpoint        *DefaultServiceEndpointLikeEndpoint `json:"likeEndpoint,omitempty"`
}

type DefaultServiceEndpointLikeEndpoint struct {
	Status  Status                `json:"status"`
	Target  AddToPlaylistEndpoint `json:"target"`
	Actions []LikeEndpointAction  `json:"actions,omitempty"`
}

type MusicTwoRowItemRendererNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *WatchEndpoint         `json:"watchEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
	BrowseEndpoint        *PurpleBrowseEndpoint  `json:"browseEndpoint,omitempty"`
}

type PurpleBrowseEndpoint struct {
	BrowseID                              string                                `json:"browseId"`
	BrowseEndpointContextSupportedConfigs BrowseEndpointContextSupportedConfigs `json:"browseEndpointContextSupportedConfigs"`
	Params                                *string                               `json:"params,omitempty"`
}

type Subtitle struct {
	Runs []SubtitleRun `json:"runs"`
}

type SubtitleRun struct {
	Text               string `json:"text"`
	NavigationEndpoint *OnTap `json:"navigationEndpoint,omitempty"`
}

type OnTap struct {
	ClickTrackingParams string              `json:"clickTrackingParams"`
	BrowseEndpoint      OnTapBrowseEndpoint `json:"browseEndpoint"`
}

type ThumbnailOverlay struct {
	MusicItemThumbnailOverlayRenderer ThumbnailOverlayMusicItemThumbnailOverlayRenderer `json:"musicItemThumbnailOverlayRenderer"`
}

type ThumbnailOverlayMusicItemThumbnailOverlayRenderer struct {
	Background      MusicItemThumbnailOverlayRendererBackground `json:"background"`
	Content         FluffyContent                               `json:"content"`
	ContentPosition ContentPosition                             `json:"contentPosition"`
	DisplayStyle    DisplayStyle                                `json:"displayStyle"`
}

type FluffyContent struct {
	MusicPlayButtonRenderer FluffyMusicPlayButtonRenderer `json:"musicPlayButtonRenderer"`
}

type FluffyMusicPlayButtonRenderer struct {
	PlayNavigationEndpoint FluffyPlayNavigationEndpoint `json:"playNavigationEndpoint"`
	TrackingParams         string                       `json:"trackingParams"`
	PlayIcon               Icon                         `json:"playIcon"`
	PauseIcon              Icon                         `json:"pauseIcon"`
	IconColor              int64                        `json:"iconColor"`
	BackgroundColor        int64                        `json:"backgroundColor"`
	ActiveBackgroundColor  int64                        `json:"activeBackgroundColor"`
	LoadingIndicatorColor  int64                        `json:"loadingIndicatorColor"`
	PlayingIcon            Icon                         `json:"playingIcon"`
	IconLoadingColor       int64                        `json:"iconLoadingColor"`
	ActiveScaleFactor      float64                      `json:"activeScaleFactor"`
	ButtonSize             FluffyButtonSize             `json:"buttonSize"`
	RippleTarget           RippleTarget                 `json:"rippleTarget"`
	AccessibilityPlayData  Accessibility                `json:"accessibilityPlayData"`
	AccessibilityPauseData Accessibility                `json:"accessibilityPauseData"`
}

type FluffyPlayNavigationEndpoint struct {
	ClickTrackingParams   string                 `json:"clickTrackingParams"`
	WatchEndpoint         *WatchEndpoint         `json:"watchEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint `json:"watchPlaylistEndpoint,omitempty"`
}

type MusicTwoRowItemRendererTitle struct {
	Runs []FluffyRun `json:"runs"`
}

type FluffyRun struct {
	Text               string                    `json:"text"`
	NavigationEndpoint *FluffyNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type FluffyNavigationEndpoint struct {
	ClickTrackingParams string               `json:"clickTrackingParams"`
	BrowseEndpoint      PurpleBrowseEndpoint `json:"browseEndpoint"`
}

type MusicCarouselShelfRendererHeader struct {
	MusicCarouselShelfBasicHeaderRenderer MusicCarouselShelfBasicHeaderRenderer `json:"musicCarouselShelfBasicHeaderRenderer"`
}

type MusicCarouselShelfBasicHeaderRenderer struct {
	Title             MusicCarouselShelfBasicHeaderRendererTitle      `json:"title"`
	Strapline         *Strapline                                      `json:"strapline,omitempty"`
	AccessibilityData Accessibility                                   `json:"accessibilityData"`
	HeaderStyle       string                                          `json:"headerStyle"`
	MoreContentButton *MoreContentButton                              `json:"moreContentButton,omitempty"`
	Thumbnail         *MusicCarouselShelfBasicHeaderRendererThumbnail `json:"thumbnail,omitempty"`
	TrackingParams    string                                          `json:"trackingParams"`
}

type MoreContentButton struct {
	ButtonRenderer MoreContentButtonButtonRenderer `json:"buttonRenderer"`
}

type MoreContentButtonButtonRenderer struct {
	Style              string                           `json:"style"`
	Text               Strapline                        `json:"text"`
	NavigationEndpoint ButtonRendererNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                           `json:"trackingParams"`
	AccessibilityData  Accessibility                    `json:"accessibilityData"`
}

type ButtonRendererNavigationEndpoint struct {
	ClickTrackingParams   string                  `json:"clickTrackingParams"`
	BrowseEndpoint        *EndpointBrowseEndpoint `json:"browseEndpoint,omitempty"`
	WatchPlaylistEndpoint *WatchPlaylistEndpoint  `json:"watchPlaylistEndpoint,omitempty"`
}

type EndpointBrowseEndpoint struct {
	BrowseID string `json:"browseId"`
}

type MusicCarouselShelfBasicHeaderRendererThumbnail struct {
	MusicThumbnailRenderer PurpleMusicThumbnailRenderer `json:"musicThumbnailRenderer"`
}

type PurpleMusicThumbnailRenderer struct {
	Thumbnail         MusicThumbnailRendererThumbnail `json:"thumbnail"`
	ThumbnailCrop     string                          `json:"thumbnailCrop"`
	ThumbnailScale    ThumbnailScale                  `json:"thumbnailScale"`
	TrackingParams    string                          `json:"trackingParams"`
	AccessibilityData Accessibility                   `json:"accessibilityData"`
	OnTap             OnTap                           `json:"onTap"`
	TargetID          string                          `json:"targetId"`
}

type MusicCarouselShelfBasicHeaderRendererTitle struct {
	Runs []TentacledRun `json:"runs"`
}

type TentacledRun struct {
	Text               string    `json:"text"`
	NavigationEndpoint *Endpoint `json:"navigationEndpoint,omitempty"`
}

type Endpoint struct {
	ClickTrackingParams string                 `json:"clickTrackingParams"`
	BrowseEndpoint      EndpointBrowseEndpoint `json:"browseEndpoint"`
}

type Continuation struct {
	NextContinuationData NextContinuationData `json:"nextContinuationData"`
}

type NextContinuationData struct {
	Continuation        string `json:"continuation"`
	ClickTrackingParams string `json:"clickTrackingParams"`
}

type SectionListRendererHeader struct {
	ChipCloudRenderer ChipCloudRenderer `json:"chipCloudRenderer"`
}

type ChipCloudRenderer struct {
	Chips                []Chip `json:"chips"`
	TrackingParams       string `json:"trackingParams"`
	HorizontalScrollable bool   `json:"horizontalScrollable"`
}

type Chip struct {
	ChipCloudChipRenderer ChipCloudChipRenderer `json:"chipCloudChipRenderer"`
}

type ChipCloudChipRenderer struct {
	Style               Style              `json:"style"`
	Text                Strapline          `json:"text"`
	NavigationEndpoint  NavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams      string             `json:"trackingParams"`
	IsSelected          bool               `json:"isSelected"`
	OnDeselectedCommand NavigationEndpoint `json:"onDeselectedCommand"`
	TargetID            *string            `json:"targetId,omitempty"`
}

type NavigationEndpoint struct {
	ClickTrackingParams string                            `json:"clickTrackingParams"`
	BrowseEndpoint      OnDeselectedCommandBrowseEndpoint `json:"browseEndpoint"`
}

type OnDeselectedCommandBrowseEndpoint struct {
	BrowseID TabIdentifier `json:"browseId"`
	Params   string        `json:"params"`
}

type Style struct {
	StyleType StyleType `json:"styleType"`
}

type ResponseContext struct {
	ServiceTrackingParams []ServiceTrackingParam `json:"serviceTrackingParams"`
	MaxAgeSeconds         int64                  `json:"maxAgeSeconds"`
}

type ServiceTrackingParam struct {
	Service string  `json:"service"`
	Params  []Param `json:"params"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ThumbnailCrop string

const (
	MusicThumbnailCropUnspecified ThumbnailCrop = "MUSIC_THUMBNAIL_CROP_UNSPECIFIED"
)

type ThumbnailScale string

const (
	MusicThumbnailScaleAspectFill  ThumbnailScale = "MUSIC_THUMBNAIL_SCALE_ASPECT_FILL"
	MusicThumbnailScaleAspectFit   ThumbnailScale = "MUSIC_THUMBNAIL_SCALE_ASPECT_FIT"
	MusicThumbnailScaleUnspecified ThumbnailScale = "MUSIC_THUMBNAIL_SCALE_UNSPECIFIED"
)

type IconType string

const (
	AddToPlaylist      IconType = "ADD_TO_PLAYLIST"
	AddToRemoteQueue   IconType = "ADD_TO_REMOTE_QUEUE"
	Album              IconType = "ALBUM"
	Artist             IconType = "ARTIST"
	Favorite           IconType = "FAVORITE"
	LibraryAdd         IconType = "LIBRARY_ADD"
	LibrarySaved       IconType = "LIBRARY_SAVED"
	Mix                IconType = "MIX"
	MusicExplicitBadge IconType = "MUSIC_EXPLICIT_BADGE"
	MusicShuffle       IconType = "MUSIC_SHUFFLE"
	Pause              IconType = "PAUSE"
	PlayArrow          IconType = "PLAY_ARROW"
	QueuePlayNext      IconType = "QUEUE_PLAY_NEXT"
	Share              IconType = "SHARE"
	TabHome            IconType = "TAB_HOME"
	Unfavorite         IconType = "UNFAVORITE"
	VolumeUp           IconType = "VOLUME_UP"
)

type DisplayPriority string

const (
	MusicResponsiveListItemColumnDisplayPriorityHigh   DisplayPriority = "MUSIC_RESPONSIVE_LIST_ITEM_COLUMN_DISPLAY_PRIORITY_HIGH"
	MusicResponsiveListItemColumnDisplayPriorityMedium DisplayPriority = "MUSIC_RESPONSIVE_LIST_ITEM_COLUMN_DISPLAY_PRIORITY_MEDIUM"
)

type PageType string

const (
	MusicPageTypeAlbum       PageType = "MUSIC_PAGE_TYPE_ALBUM"
	MusicPageTypeArtist      PageType = "MUSIC_PAGE_TYPE_ARTIST"
	MusicPageTypePlaylist    PageType = "MUSIC_PAGE_TYPE_PLAYLIST"
	MusicPageTypeUserChannel PageType = "MUSIC_PAGE_TYPE_USER_CHANNEL"
)

type Params string

const (
	WAEB Params = "wAEB"
)

type MusicVideoType string

const (
	MusicVideoTypeAtv MusicVideoType = "MUSIC_VIDEO_TYPE_ATV"
)

type ItemHeight string

const (
	MusicResponsiveListItemHeightMediumCompact ItemHeight = "MUSIC_RESPONSIVE_LIST_ITEM_HEIGHT_MEDIUM_COMPACT"
)

type Status string

const (
	Dislike     Status = "DISLIKE"
	Indifferent Status = "INDIFFERENT"
	Like        Status = "LIKE"
)

type SharePanelType string

const (
	SharePanelTypeUnifiedSharePanel SharePanelType = "SHARE_PANEL_TYPE_UNIFIED_SHARE_PANEL"
)

type QueueInsertPosition string

const (
	InsertAfterCurrentVideo QueueInsertPosition = "INSERT_AFTER_CURRENT_VIDEO"
	InsertAtEnd             QueueInsertPosition = "INSERT_AT_END"
)

type LibraryStatus string

const (
	MusicLibraryStatusInLibrary LibraryStatus = "MUSIC_LIBRARY_STATUS_IN_LIBRARY"
)

type PurpleButtonSize string

const (
	MusicPlayButtonSizeSmall PurpleButtonSize = "MUSIC_PLAY_BUTTON_SIZE_SMALL"
)

type RippleTarget string

const (
	MusicPlayButtonRippleTargetAncestor RippleTarget = "MUSIC_PLAY_BUTTON_RIPPLE_TARGET_ANCESTOR"
	MusicPlayButtonRippleTargetSelf     RippleTarget = "MUSIC_PLAY_BUTTON_RIPPLE_TARGET_SELF"
)

type ContentPosition string

const (
	MusicItemThumbnailOverlayContentPositionBottomRight ContentPosition = "MUSIC_ITEM_THUMBNAIL_OVERLAY_CONTENT_POSITION_BOTTOM_RIGHT"
	MusicItemThumbnailOverlayContentPositionCentered    ContentPosition = "MUSIC_ITEM_THUMBNAIL_OVERLAY_CONTENT_POSITION_CENTERED"
)

type DisplayStyle string

const (
	MusicItemThumbnailOverlayDisplayStyleHover      DisplayStyle = "MUSIC_ITEM_THUMBNAIL_OVERLAY_DISPLAY_STYLE_HOVER"
	MusicItemThumbnailOverlayDisplayStylePersistent DisplayStyle = "MUSIC_ITEM_THUMBNAIL_OVERLAY_DISPLAY_STYLE_PERSISTENT"
)

type AspectRatio string

const (
	MusicTwoRowItemThumbnailAspectRatioSquare AspectRatio = "MUSIC_TWO_ROW_ITEM_THUMBNAIL_ASPECT_RATIO_SQUARE"
)

type FluffyButtonSize string

const (
	MusicPlayButtonSizeHuge   FluffyButtonSize = "MUSIC_PLAY_BUTTON_SIZE_HUGE"
	MusicPlayButtonSizeMedium FluffyButtonSize = "MUSIC_PLAY_BUTTON_SIZE_MEDIUM"
)

type TabIdentifier string

const (
	FEmusicHome TabIdentifier = "FEmusic_home"
)

type StyleType string

const (
	StyleLargeTranslucentAndSelectedWhite StyleType = "STYLE_LARGE_TRANSLUCENT_AND_SELECTED_WHITE"
)
