package client


type Initiator struct {
	Iqn     string   `json:"iqn"`
	Nqn     string   `json:"nqn"`
	Portal  string   `json:"portal"`
	Wwn     string   `json:"wwn"`
}

type Target struct {
	Name    string   `json:"name"`
	Iqn     string   `json:"iqn"`
	Nqn     string   `json:"nqn"`
	Portal  string   `json:"portal"`
	Wwn     string   `json:"wwn"`
        Failover string  `json:"failover"`
}

type HostBalance struct {
	Name                       string     `json:"name"`
	Op_count                   int     `json:"op_count"`
	Fraction_relative_to_max   float64     `json:"fraction_relative_to_max"`
	Initiator                  Initiator     `json:"initiator"`
	Target                     Target     `json:"target"`
	Time                       int     `json:"time"`
}

type HostsBalanceList struct {
        ContinuationToken    string   `json:"continuation_token"`
        TotalItemCount       int      `json:"total_item_count"`
        MoreItemsRemaining   bool     `json:"more_items_remaining"`
        Items                []HostBalance `json:"items"`
}

func (fa *FAClient) GetHostsBalance() *HostsBalanceList {
        result := new(HostsBalanceList)
        res, err := fa.RestClient.R().
                SetResult(&result).
                Get("/hosts/performance/balance")

        if err != nil {
                fa.Error = err
        }
        if res.StatusCode() == 401 {
                fa.RefreshSession()
        }
        res, err = fa.RestClient.R().
                SetResult(&result).
                Get("/hosts/performance/balance")
        if err != nil {
                fa.Error = err
        }

        return result
}
