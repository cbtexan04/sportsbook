package dtos

import "time"

type DraftKingsResponse struct {
	EventGroup struct {
		EventGroupID         int    `json:"eventGroupId"`
		ProviderEventGroupID string `json:"providerEventGroupId"`
		ProviderID           int    `json:"providerId"`
		DisplayGroupID       int    `json:"displayGroupId"`
		Name                 string `json:"name"`
		OfferCategories      []struct {
			OfferCategoryID             int    `json:"offerCategoryId"`
			Name                        string `json:"name"`
			OfferSubcategoryDescriptors []struct {
				SubcategoryID    int    `json:"subcategoryId"`
				Name             string `json:"name"`
				OfferSubcategory struct {
					Name          string `json:"name"`
					SubcategoryID int    `json:"subcategoryId"`
					ComponentID   int    `json:"componentId"`
					Offers        [][]Offers `json:"offers"`
				} `json:"offerSubcategory,omitempty"`
			} `json:"offerSubcategoryDescriptors,omitempty"`
		} `json:"offerCategories"`
		Events []struct {
			EventID         int       `json:"eventId"`
			DisplayGroupID  int       `json:"displayGroupId"`
			EventGroupID    int       `json:"eventGroupId"`
			EventGroupName  string    `json:"eventGroupName"`
			ProviderEventID string    `json:"providerEventId"`
			ProviderID      int       `json:"providerId"`
			Name            string    `json:"name"`
			StartDate       time.Time `json:"startDate"`
			TeamName1       string    `json:"teamName1"`
			TeamName2       string    `json:"teamName2"`
			TeamShortName1  string    `json:"teamShortName1"`
			TeamShortName2  string    `json:"teamShortName2"`
			EventStatus     struct {
				State           string `json:"state"`
				IsClockDisabled bool   `json:"isClockDisabled"`
				Minute          int    `json:"minute"`
				Second          int    `json:"second"`
				IsClockRunning  bool   `json:"isClockRunning"`
			} `json:"eventStatus"`
			EventScorecard struct {
				ScorecardComponentID int `json:"scorecardComponentId"`
			} `json:"eventScorecard"`
			MediaList []struct {
				EventID           int       `json:"eventId"`
				MediaProviderName string    `json:"mediaProviderName"`
				MediaID           string    `json:"mediaId"`
				MediaTypeName     string    `json:"mediaTypeName"`
				UpdatedAt         time.Time `json:"updatedAt"`
			} `json:"mediaList"`
			LiveBettingOffered       bool `json:"liveBettingOffered"`
			LiveBettingEnabled       bool `json:"liveBettingEnabled"`
			IsSameGameParlayEligible bool `json:"isSameGameParlayEligible,omitempty"`
		} `json:"events"`
	} `json:"eventGroup"`
	Promotions []struct {
		PromotionID                int       `json:"promotionId"`
		IsArchived                 bool      `json:"isArchived"`
		SortOrder                  int       `json:"sortOrder"`
		ShowInCarousel             bool      `json:"showInCarousel"`
		ShowOnPromoPage            bool      `json:"showOnPromoPage"`
		IsFeatured                 bool      `json:"isFeatured"`
		TaggedEntityType           int       `json:"taggedEntityType,omitempty"`
		ProviderID                 int       `json:"providerId,omitempty"`
		EventGroupID               int       `json:"eventGroupId,omitempty"`
		StartDate                  time.Time `json:"startDate"`
		ProductID                  int       `json:"productId"`
		JurisdictionIds            []int     `json:"jurisdictionIds"`
		IsOptedIn                  bool      `json:"isOptedIn"`
		PromotionHeadline          string    `json:"promotionHeadline"`
		PromotionDescription       string    `json:"promotionDescription"`
		CtaLink                    string    `json:"ctaLink"`
		CtaDeepLink                string    `json:"ctaDeepLink"`
		ExpirationDate             time.Time `json:"expirationDate"`
		SiteExperience             []string  `json:"siteExperience"`
		ProductName                string    `json:"productName"`
		OptIn                      bool      `json:"optIn"`
		IsTargeted                 bool      `json:"isTargeted"`
		SortAheadOfTargetedPromos  bool      `json:"sortAheadOfTargetedPromos"`
		IsOfferCardImagePromotion  bool      `json:"isOfferCardImagePromotion"`
		ShouldImageCardAutoRefresh bool      `json:"shouldImageCardAutoRefresh"`
		IsStraplinePromotion       bool      `json:"isStraplinePromotion"`
		WebStraplineZone           string    `json:"webStraplineZone,omitempty"`
		NativeStraplineZone        string    `json:"nativeStraplineZone,omitempty"`
		WebStraplineImage          string    `json:"webStraplineImage,omitempty"`
		NativeStraplineImage       string    `json:"nativeStraplineImage,omitempty"`
		IsFTDOnlySportsbook        bool      `json:"isFTDOnlySportsbook"`
		IsFTDOnlyCasino            bool      `json:"isFTDOnlyCasino"`
		ExternalCampaignID         int       `json:"externalCampaignId,omitempty"`
		DisplayGroupID             int       `json:"displayGroupId,omitempty"`
		DestinationURL             string    `json:"destinationUrl,omitempty"`
		Deeplink                   string    `json:"deeplink,omitempty"`
		PromotionImage             string    `json:"promotionImage,omitempty"`
		PromotionSubHeadline       string    `json:"promotionSubHeadline,omitempty"`
		CtaText                    string    `json:"ctaText,omitempty"`
		PromotionTermsModalImage   string    `json:"promotionTermsModalImage,omitempty"`
		Terms                      string    `json:"terms,omitempty"`
		PromotionTypeName          string    `json:"promotionTypeName,omitempty"`
		PromotionOfferCardImage    string    `json:"promotionOfferCardImage,omitempty"`
		ImageCardHasTerms          bool      `json:"imageCardHasTerms,omitempty"`
		EventID                    int       `json:"eventId,omitempty"`
	} `json:"promotions"`
}

type Offers struct {
	ProviderOfferID       string `json:"providerOfferId"`
	ProviderID            int    `json:"providerId"`
	ProviderEventID       string `json:"providerEventId"`
	ProviderEventGroupID  string `json:"providerEventGroupId"`
	Label                 string `json:"label"`
	IsSuspended           bool   `json:"isSuspended"`
	IsOpen                bool   `json:"isOpen"`
	OfferSubcategoryID    int    `json:"offerSubcategoryId"`
	IsSubcategoryFeatured bool   `json:"isSubcategoryFeatured"`
	BetOfferTypeID        int    `json:"betOfferTypeId"`
	DisplayGroupID        int    `json:"displayGroupId"`
	EventGroupID          int    `json:"eventGroupId"`
	ProviderCriterionID   string `json:"providerCriterionId"`
	Outcomes              []struct {
		ProviderOutcomeID  string  `json:"providerOutcomeId"`
		ProviderID         int     `json:"providerId"`
		ProviderOfferID    string  `json:"providerOfferId"`
		Label              string  `json:"label"`
		OddsAmerican       string  `json:"oddsAmerican"`
		OddsDecimal        float64 `json:"oddsDecimal"`
		OddsDecimalDisplay string  `json:"oddsDecimalDisplay"`
		OddsFractional     string  `json:"oddsFractional"`
		Line               float64 `json:"line"`
		Main               bool    `json:"main"`
	} `json:"outcomes"`
	OfferUpdateState string `json:"offerUpdateState"`
	OfferSequence    int    `json:"offerSequence"`
	Source           string `json:"source"`
}