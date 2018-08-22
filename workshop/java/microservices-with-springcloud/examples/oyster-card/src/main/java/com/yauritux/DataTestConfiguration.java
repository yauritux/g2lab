package com.yauritux;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.Fare;
import com.yauritux.model.entity.Station;
import com.yauritux.model.entity.Zone;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Configuration
public class DataTestConfiguration {

	private static final String[] ZONES = { "1", "2", "3" };
	
	private static final String[] STATION_NAMES =  {
			"Holborn", "Earl's Court", "Wimbledon", "Hammersmith"
	};
	
	private Map<String, Station> stations = new HashMap<>();
	private Map<String, Zone> zones = new HashMap<>();
	private Map<String, Fare> fares = new HashMap<>();
	
	private Card card;
		
	DataTestConfiguration() {
		
		card = new Card();
		card.setOwner("Yauri Attamimi");
		card.setSerialNo("1234567890");
		card.setBalance(BigDecimal.ZERO); // initial balance
		
		Station holbornStation = new Station();
		holbornStation.setName(STATION_NAMES[0]);
		Station earlsCourtStation = new Station();
		earlsCourtStation.setName(STATION_NAMES[1]);
		Station wimbledonStation = new Station();
		wimbledonStation.setName(STATION_NAMES[2]);
		Station hammersmithStation = new Station();
		hammersmithStation.setName(STATION_NAMES[3]);

		stations.put(STATION_NAMES[0], holbornStation);
		stations.put(STATION_NAMES[1], earlsCourtStation);
		stations.put(STATION_NAMES[2], wimbledonStation);
		stations.put(STATION_NAMES[3], hammersmithStation);
		
		Zone zone1 = new Zone();
		zone1.setZoneName(ZONES[0]);
		Zone zone2 = new Zone();
		zone2.setZoneName(ZONES[1]);
		Zone zone3 = new Zone();
		zone3.setZoneName(ZONES[2]);
		
		stations.get(STATION_NAMES[0]).addZone(zone1); // holborn
		stations.get(STATION_NAMES[1]).addZone(zone1); // earls court
		stations.get(STATION_NAMES[1]).addZone(zone2); // earls court
		stations.get(STATION_NAMES[2]).addZone(zone3); // wimbledon
		stations.get(STATION_NAMES[3]).addZone(zone2); // hammersmith
		
		zones.put(ZONES[0], zone1);
		zones.put(ZONES[1], zone2);
		zones.put(ZONES[2], zone3);
		
		// Anywhere in Zone 1
		Fare fare1 = new Fare();
		fare1.setJourney("Tube#1");
		fare1.setJourneyType(JourneyType.TUBE);
		fare1.setFare(BigDecimal.valueOf(2.50));
		zones.get(ZONES[0]).addFare(fare1);
		
		fares.put(fare1.getJourney(), fare1);
		
		// Any one Zone outside zone 1
		Fare fare2 = new Fare();
		fare2.setJourney("Tube#2");
		fare2.setJourneyType(JourneyType.TUBE);
		fare2.setFare(BigDecimal.valueOf(2.00));
		zones.get(ZONES[1]).addFare(fare2);
		
		fares.put(fare2.getJourney(), fare2);
		
		Fare fare3 = new Fare();
		fare3.setJourney("Tube#3");
		fare3.setJourneyType(JourneyType.TUBE);
		fare3.setFare(BigDecimal.valueOf(2.00));
		zones.get(ZONES[2]).addFare(fare3);
		
		fares.put(fare3.getJourney(), fare3);
		
		// Any two zones including zone 1
		Fare fare4 = new Fare();
		fare4.setJourney("Tube#4");
		fare4.setJourneyType(JourneyType.TUBE);
		fare4.setFare(BigDecimal.valueOf(3.00));
		zones.get(ZONES[0]).addFare(fare4);
		zones.get(ZONES[1]).addFare(fare4);
		
		fares.put(fare4.getJourney(), fare4);
		
		Fare fare5 = new Fare();
		fare5.setJourney("Tube#5");
		fare5.setJourneyType(JourneyType.TUBE);
		fare5.setFare(BigDecimal.valueOf(3.00));
		zones.get(ZONES[1]).addFare(fare5);
		zones.get(ZONES[2]).addFare(fare5);
		
		fares.put(fare5.getJourney(), fare5);
		
		Fare fare6 = new Fare();
		fare6.setJourney("Tube#6");
		fare6.setJourneyType(JourneyType.TUBE);
		fare6.setFare(BigDecimal.valueOf(3.00));
		zones.get(ZONES[0]).addFare(fare6);
		zones.get(ZONES[2]).addFare(fare6);
		
		fares.put(fare6.getJourney(), fare6);
		
		// Any two zones excluding zone 1
		Fare fare7 = new Fare();
		fare7.setJourney("Tube#7");
		fare7.setJourneyType(JourneyType.TUBE);
		fare7.setFare(BigDecimal.valueOf(2.25));
		zones.get(ZONES[1]).addFare(fare7);
		zones.get(ZONES[2]).addFare(fare7);
		
		fares.put(fare7.getJourney(), fare7);
		
		// Any three zones
		Fare fare8 = new Fare();
		fare8.setJourney("Tube#8");
		fare8.setJourneyType(JourneyType.TUBE);
		fare8.setFare(BigDecimal.valueOf(3.20));
		zones.get(ZONES[0]).addFare(fare8);
		zones.get(ZONES[1]).addFare(fare8);
		zones.get(ZONES[2]).addFare(fare8);
		
		fares.put(fare8.getJourney(), fare8);
		
		// Any bus journey
		Fare fare9 = new Fare();
		fare9.setJourney("BusFare");
		fare9.setJourneyType(JourneyType.BUS);
		fare9.setFare(BigDecimal.valueOf(1.80));
		
		fares.put(fare9.getJourney(), fare9);			
	}
	
	@Bean
	public Card getCard() {
		return card;
	}
	
	@Bean
	public Map<String, Station> stations() {		
		return stations;
	}
	
	@Bean
	public Map<String, Zone> zones() {		
		return zones;
	}
	
	@Bean
	public Map<String, Fare> fares() {		
		return fares;
	}
}
