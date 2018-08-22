package com.yauritux.repository;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.assertNull;

import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.orm.jpa.DataJpaTest;
import org.springframework.boot.test.autoconfigure.orm.jpa.TestEntityManager;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import com.yauritux.Application;
import com.yauritux.model.entity.Station;
import com.yauritux.model.entity.Zone;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@RunWith(SpringRunner.class)
@SpringBootTest(classes = Application.class)
@DataJpaTest
public class StationRepositoryTest {
	
	private static final String[] STATION_NAMES =  {
			"Holborn", "Earl's Court", "Wimbledon", "Hammersmith"
	};
	
	private static final String[] ZONE_NAMES = {
			"1", "2", "3"
	};

	@Autowired
	private TestEntityManager em;
	
	@Autowired
	private StationRepository stationRepository;
	
	@Before
	public void setup() {
		Station holbornStation = new Station();
		holbornStation.setName(STATION_NAMES[0]);
		Station earlsCourtStation = new Station();
		earlsCourtStation.setName(STATION_NAMES[1]);
		Station wimbledonStation = new Station();
		wimbledonStation.setName(STATION_NAMES[2]);
		Station hammersmithStation = new Station();
		hammersmithStation.setName(STATION_NAMES[3]);
		
		Zone zone1 = new Zone();
		zone1.setZoneName(ZONE_NAMES[0]);
		Zone zone2 = new Zone();
		zone2.setZoneName(ZONE_NAMES[1]);
		Zone zone3 = new Zone();
		zone3.setZoneName(ZONE_NAMES[2]);
		
		holbornStation.addZone(zone1);
		earlsCourtStation.addZone(zone1);
		earlsCourtStation.addZone(zone2);
		wimbledonStation.addZone(zone3);
		hammersmithStation.addZone(zone2);

		em.persist(holbornStation);
		em.persist(earlsCourtStation);
		em.persist(wimbledonStation);
		em.persist(hammersmithStation);
		em.flush();
	}
	
	@After
	public void tearDown() {
		stationRepository.delete(stationRepository.findByName(STATION_NAMES[0]));
		stationRepository.delete(stationRepository.findByName(STATION_NAMES[1]));
		stationRepository.delete(stationRepository.findByName(STATION_NAMES[2]));
		stationRepository.delete(stationRepository.findByName(STATION_NAMES[3]));
	}
	
	@Test
	public void findByName_existingStation_shouldReturnStation() {
		Station found = stationRepository.findByName(STATION_NAMES[0]);
		assertThat(found.getName()).isEqualTo(STATION_NAMES[0]);
	}
	
	@Test
	public void findByName_holbornStation_hasOneZone() {
		Station holborn = stationRepository.findByName(STATION_NAMES[0]);
		assertThat(holborn.getZones().size()).isEqualTo(1);
	}
	
	@Test
	public void findByName_nonExistingStation_shouldReturnNothing() {
		Station notFound = stationRepository.findByName("NOT_REGISTERED");
		assertNull(notFound);
	}
}
