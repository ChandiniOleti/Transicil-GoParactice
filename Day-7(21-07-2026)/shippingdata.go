package main
import "fmt"
import "os"
import "encoding/json"
import "sort"
import "time"

type Metadata struct {
	GeneratedOn  string `json:"generated_on"`
	RecordCount  int    `json:"record_count"`
	SchemaVersion string `json:"schema_version"`
}

type Origin struct {
	Country      string `json:"country"`
	City         string `json:"city"`
	WarehouseCode string `json:"warehouse_code"`
}

type Destination struct {
	Country    string `json:"country"`
	City       string `json:"city"`
	AddressLine string `json:"address_line"`
	PostalCode string `json:"postal_code"`
}

type Item struct {
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type CostBreakdown struct {
	BaseRate       float64 `json:"base_rate"`
	FuelSurcharge  float64 `json:"fuel_surcharge"`
	Insurance      float64 `json:"insurance"`
	Tax            float64 `json:"tax"`
	Total          float64 `json:"total"`
	Currency       string  `json:"currency"`
}

type TrackingEvent struct {
	Timestamp    string `json:"timestamp"`
	Location     string `json:"location"`
	StatusUpdate string `json:"status_update"`
}

type Shipment struct {
	ShipmentID        string           `json:"shipment_id"`
	OrderID           string           `json:"order_id"`
	Carrier           string           `json:"carrier"`
	Status            string           `json:"status"`
	IsInternational   bool             `json:"is_international"`
	Origin            Origin           `json:"origin"`
	Destination       Destination      `json:"destination"`
	Items             []Item           `json:"items"`
	CostBreakdown     CostBreakdown    `json:"cost_breakdown"`
	WeightKg          float64          `json:"weight_kg"`
	CreatedAt         string           `json:"created_at"`
	EstimatedDelivery string           `json:"estimated_delivery"`
	TrackingEvents    []TrackingEvent  `json:"tracking_events"`
}

type ShipmentData struct {
	Metadata  Metadata   `json:"metadata"`
	Shipments []Shipment `json:"shipments"`
}
func main(){
	data,err:=os.ReadFile("shipping_data.json")
	if err!=nil{
		fmt.Println(err)
		return
	}
	var shipdata ShipmentData
	err=json.Unmarshal(data,&shipdata)
	if err!=nil{
		fmt.Println(err)
		return
	}
	// fmt.Printf(shipdata)

	//====Filtering + Sorting: Find all shipments where status is "Delayed" or "Customs Hold", and sort them by estimated_delivery in ascending order.
	var filtered []Shipment
	for _,shipmentv:=range shipdata.Shipments{
		if shipmentv.Status=="Delayed" || shipmentv.Status=="Customs Hold"{
			filtered=append(filtered,shipmentv)
		}
		// fmt.Println(filtered)
	}

	//=====sorting them by estimated
	sort.Slice(filtered,func(i,j int) bool{
		return filtered[i].EstimatedDelivery<filtered[j].EstimatedDelivery
	})
	fmt.Println("After sorting")
	for _,sortdata:=range filtered{

		fmt.Println(sortdata.EstimatedDelivery," ",sortdata.Status)
	}

	//============Aggregation: Calculate the total revenue (cost_breakdown.total) grouped by carrier. Which carrier generated the most revenue?

	revenue:=make(map[string]float64)
	for _,shipmentre:=range shipdata.Shipments{
		revenue[shipmentre.Carrier]+=shipmentre.CostBreakdown.Total
	}
	fmt.Println(revenue)
	
	maxrevenue:=0.0
	bestCarrier:=""
	for carrier,total:=range revenue{
		if total>maxrevenue{
			maxrevenue=total
			bestCarrier=carrier
		}
	}
	fmt.Println(bestCarrier)
	fmt.Println(maxrevenue)
	

	//==========Nested Array Traversal: For each shipment, find the number of hours between the first and last event in tracking_events. Return the shipment with the longest total transit time.

	var maxHours float64
	var maxShipment Shipment

	for _, shipment := range shipdata.Shipments {

		if len(shipment.TrackingEvents) == 0 {
			continue
		}

		first := shipment.TrackingEvents[0]
		last := shipment.TrackingEvents[len(shipment.TrackingEvents)-1]

		startTime, err := time.Parse(time.RFC3339, first.Timestamp)//Convert timestamps into time
		if err != nil {
			continue
		}

		endTime, err := time.Parse(time.RFC3339, last.Timestamp)//Convert timestamps into time using time.parse
		if err != nil {
			continue
		}

		hours := endTime.Sub(startTime).Hours()

		if hours > maxHours {
			maxHours = hours
			maxShipment = shipment
		}
	}

	fmt.Println("Shipment ID :", maxShipment.ShipmentID)
	fmt.Println("Carrier     :", maxShipment.Carrier)
	fmt.Printf("Transit Time: %.2f hours\n", maxHours)


	//===========Cross-field Logic: Identify all is_international shipments where weight_kg > 20 and cost_breakdown.insurance == 0. These represent heavy international shipments with no insurance — flag them as a "risk report."
	fmt.Println("------ RISK REPORT ------")
	for _,shipment:=range shipdata.Shipments{
		if shipment.IsInternational && shipment.WeightKg>20 && shipment.CostBreakdown.Insurance==0{
			// fmt.Println(shipment)
			fmt.Println("Shipment ID :", shipment.ShipmentID)
			fmt.Println("Carrier     :", shipment.Carrier)
			fmt.Println("Weight      :", shipment.WeightKg)
			fmt.Println("Insurance   :", shipment.CostBreakdown.Insurance)
		
		}
	}


	//==========Multi-level Grouping: Build a summary object grouping shipments by destination.country, and within each country, count shipments per status.

	summary := make(map[string]map[string]int)

	for _, shipment := range shipdata.Shipments {

		country := shipment.Destination.Country
		status := shipment.Status

		if summary[country] == nil {
			summary[country] = make(map[string]int)
		}

		summary[country][status]++
		
	}
	result, err := json.MarshalIndent(summary, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))
}